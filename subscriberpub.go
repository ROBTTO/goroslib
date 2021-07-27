package goroslib

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"time"

	"github.com/ROBTTO/goroslib/pkg/apislave"
	"github.com/ROBTTO/goroslib/pkg/protocommon"
	"github.com/ROBTTO/goroslib/pkg/prototcp"
	"github.com/ROBTTO/goroslib/pkg/protoudp"
)

var errSubscriberPubTerminate = errors.New("subscriberPublisher terminated")

type subscriberPublisher struct {
	sub     *Subscriber
	address string

	ctx       context.Context
	ctxCancel func()
	udpAddr   *net.UDPAddr
	udpID     uint32

	// in
	udpFrame chan *protoudp.Frame
}

func newSubscriberPublisher(sub *Subscriber, address string) {
	ctx, ctxCancel := context.WithCancel(sub.ctx)

	sp := &subscriberPublisher{
		sub:       sub,
		address:   address,
		ctx:       ctx,
		ctxCancel: ctxCancel,
	}

	sub.publishers[address] = sp

	sub.publishersWg.Add(1)
	go sp.run()
}

func (sp *subscriberPublisher) close() {
	delete(sp.sub.publishers, sp.address)
	sp.ctxCancel()
}

func (sp *subscriberPublisher) run() {
	defer sp.sub.publishersWg.Done()

	for {
		ok := func() bool {
			err := sp.runInner()
			if err == errSubscriberPubTerminate {
				return false
			}

			t := time.NewTimer(5 * time.Second)
			defer t.Stop()

			select {
			case <-t.C:
				return true

			case <-sp.ctx.Done():
				return false
			}
		}()
		if !ok {
			break
		}
	}

	sp.ctxCancel()
}

func (sp *subscriberPublisher) runInner() error {
	xcs := apislave.NewClient(sp.address, sp.sub.conf.Node.absoluteName())

	subDone := make(chan struct{}, 1)
	var proto []interface{}
	var err error

	go func() {
		defer close(subDone)

		protocols := func() [][]interface{} {
			if sp.sub.conf.Protocol == TCP {
				return [][]interface{}{{"TCPROS"}}
			}

			nodeIP, _, _ := net.SplitHostPort(sp.sub.conf.Node.nodeAddr.String())
			return [][]interface{}{{
				"UDPROS",
				func() []byte {
					var buf bytes.Buffer
					protocommon.HeaderEncode(&buf, &protoudp.HeaderSubscriber{
						Callerid: sp.sub.conf.Node.absoluteName(),
						Md5sum:   sp.sub.msgMd5,
						Topic:    sp.sub.conf.Node.absoluteTopicName(sp.sub.conf.Topic),
						Type:     sp.sub.msgType,
					})
					return buf.Bytes()[4:]
				}(),
				nodeIP,
				sp.sub.conf.Node.udprosServer.Port(),
				1500,
			}}
		}()
		proto, err = xcs.RequestTopic(sp.sub.conf.Node.absoluteTopicName(sp.sub.conf.Topic), protocols)
	}()

	select {
	case <-subDone:
	case <-sp.ctx.Done():
		<-subDone
		return errSubscriberPubTerminate
	}

	if err != nil {
		return err
	}

	if sp.sub.conf.Protocol == TCP {
		return sp.runInnerTCP(proto)
	}
	return sp.runInnerUDP(proto)
}

func (sp *subscriberPublisher) runInnerTCP(proto []interface{}) error {
	if len(proto) != 3 {
		return fmt.Errorf("wrong protocol length")
	}

	protoName, ok := proto[0].(string)
	if !ok {
		return fmt.Errorf("wrong protoName")
	}

	protoHost, ok := proto[1].(string)
	if !ok {
		return fmt.Errorf("wrong protoHost")
	}

	protoPort, ok := proto[2].(int)
	if !ok {
		return fmt.Errorf("wrong protoPort")
	}

	if protoName != "TCPROS" {
		return fmt.Errorf("wrong protoName")
	}

	addr := net.JoinHostPort(protoHost, strconv.FormatInt(int64(protoPort), 10))

	subDone := make(chan struct{}, 1)
	var conn *prototcp.Conn
	var err error
	go func() {
		defer close(subDone)
		conn, err = prototcp.NewClient(addr)
	}()

	select {
	case <-subDone:
	case <-sp.ctx.Done():
		return errSubscriberPubTerminate
	}

	if err != nil {
		return err
	}

	defer conn.Close()

	if sp.sub.conf.EnableKeepAlive {
		conn.NetConn().(*net.TCPConn).SetKeepAlive(true)
		conn.NetConn().(*net.TCPConn).SetKeepAlivePeriod(60 * time.Second)
	}

	if sp.sub.conf.DisableNoDelay {
		err := conn.NetConn().(*net.TCPConn).SetNoDelay(false)
		if err != nil {
			return err
		}
	}

	subDone = make(chan struct{})
	var outHeader prototcp.HeaderPublisher
	go func() {
		defer close(subDone)

		err = conn.WriteHeader(&prototcp.HeaderSubscriber{
			Callerid: sp.sub.conf.Node.absoluteName(),
			Md5sum:   sp.sub.msgMd5,
			Topic:    sp.sub.conf.Node.absoluteTopicName(sp.sub.conf.Topic),
			Type:     sp.sub.msgType,
			TcpNodelay: func() int {
				if sp.sub.conf.DisableNoDelay {
					return 0
				}
				return 1
			}(),
		})
		if err != nil {
			return
		}

		var raw protocommon.HeaderRaw
		raw, err = conn.ReadHeaderRaw()
		if err != nil {
			return
		}

		if strErr, ok := raw["error"]; ok {
			err = fmt.Errorf(strErr)
			return
		}

		err = protocommon.HeaderDecode(raw, &outHeader)
	}()

	select {
	case <-subDone:
	case <-sp.ctx.Done():
		conn.Close()
		<-subDone
		return errSubscriberPubTerminate
	}

	if err != nil {
		return err
	}

	if outHeader.Topic != "" &&
		outHeader.Topic != sp.sub.conf.Node.absoluteTopicName(sp.sub.conf.Topic) {
		return fmt.Errorf("wrong topic (expected '%s', got '%s')",
			sp.sub.conf.Node.absoluteTopicName(sp.sub.conf.Topic),
			outHeader.Topic)
	}

	if outHeader.Md5sum != sp.sub.msgMd5 {
		return fmt.Errorf("wrong md5")
	}

	if sp.sub.conf.onPublisher != nil {
		sp.sub.conf.onPublisher()
	}

	subDone = make(chan struct{})
	go func() {
		defer close(subDone)

		for {
			msg := reflect.New(sp.sub.msgMsg).Interface()
			err = conn.ReadMessage(msg)
			if err != nil {
				return
			}

			if sp.sub.conf.QueueSize == 0 {
				select {
				case sp.sub.message <- msg:
				case <-sp.sub.ctx.Done():
				}
			} else {
				select {
				case sp.sub.message <- msg:
				default:
				}
			}
		}
	}()

	select {
	case <-subDone:
		return err

	case <-sp.ctx.Done():
		conn.Close()
		<-subDone
		return errSubscriberPubTerminate
	}
}

func (sp *subscriberPublisher) runInnerUDP(proto []interface{}) error {
	if len(proto) != 6 {
		return fmt.Errorf("wrong protocol length")
	}

	protoName, ok := proto[0].(string)
	if !ok {
		return fmt.Errorf("wrong protoName")
	}

	protoHost, ok := proto[1].(string)
	if !ok {
		return fmt.Errorf("wrong protoHost")
	}

	protoPort, ok := proto[2].(int)
	if !ok {
		return fmt.Errorf("wrong protoPort")
	}

	protoID, ok := proto[3].(int)
	if !ok {
		return fmt.Errorf("wrong protoID")
	}

	if protoName != "UDPROS" {
		return fmt.Errorf("wrong protoName")
	}

	// solve host and port
	addr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(protoHost, strconv.FormatInt(int64(protoPort), 10)))
	if err != nil {
		return fmt.Errorf("unable to solve host")
	}

	sp.udpAddr = addr
	sp.udpID = uint32(protoID)
	sp.udpFrame = make(chan *protoudp.Frame)

	select {
	case sp.sub.conf.Node.udpSubPublisherNew <- sp:
	case <-sp.sub.conf.Node.ctx.Done():
	}

	defer func() {
		done := make(chan struct{})
		select {
		case sp.sub.conf.Node.udpSubPublisherClose <- udpSubPublisherCloseReq{sp, done}:
			<-done
		case <-sp.sub.conf.Node.ctx.Done():
		}
		close(sp.udpFrame)
	}()

	if sp.sub.conf.onPublisher != nil {
		sp.sub.conf.onPublisher()
	}

	readerClose := make(chan struct{})
	readerDone := make(chan struct{})
	if sp.sub.conf.EnableKeepAlive {
		go func() {
			defer close(readerDone)

			curMessageID := uint8(0)

			t := time.NewTicker(60 * time.Second)
			defer t.Stop()

			for {
				select {
				case <-t.C:
					sp.sub.conf.Node.udprosServer.WriteFrame(&protoudp.Frame{
						ConnectionID: sp.udpID,
						Opcode:       protoudp.Ping,
						MessageID:    curMessageID,
					}, sp.udpAddr)
					curMessageID++

				case <-readerClose:
					return
				}
			}
		}()
	}

	var curMsg []byte
	curFieldID := 0
	curFieldCount := 0

	for {
		select {
		case frame := <-sp.udpFrame:
			switch frame.Opcode {
			case protoudp.Data0:
				curMsg = append([]byte{}, frame.Payload...)
				curFieldID = 0
				curFieldCount = int(frame.BlockID)

			case protoudp.DataN:
				if int(frame.BlockID) != (curFieldID + 1) {
					continue
				}
				curMsg = append(curMsg, frame.Payload...)
				curFieldID++
			}

			if (curFieldID + 1) == curFieldCount {
				msg := reflect.New(sp.sub.msgMsg).Interface()
				err := protocommon.MessageDecode(bytes.NewBuffer(curMsg), msg)
				if err != nil {
					continue
				}

				if sp.sub.conf.QueueSize == 0 {
					sp.sub.message <- msg
				} else {
					select {
					case sp.sub.message <- msg:
					default:
					}
				}
			}

		case <-sp.ctx.Done():
			if sp.sub.conf.EnableKeepAlive {
				close(readerClose)
				<-readerDone
			}

			return errSubscriberPubTerminate
		}
	}
}
