package goroslib

import (
	"bytes"
	"context"
	"net"

	"github.com/ROBTTO/goroslib/pkg/protocommon"
	"github.com/ROBTTO/goroslib/pkg/prototcp"
	"github.com/ROBTTO/goroslib/pkg/protoudp"
)

type publisherSubscriber struct {
	pub       *Publisher
	callerID  string
	tcpClient *prototcp.Conn
	udpAddr   *net.UDPAddr

	ctx          context.Context
	ctxCancel    func()
	curMessageID uint8
}

func newPublisherSubscriber(
	pub *Publisher,
	callerID string,
	tcpClient *prototcp.Conn,
	udpAddr *net.UDPAddr) {
	ctx, ctxCancel := context.WithCancel(pub.ctx)

	ps := &publisherSubscriber{
		pub:       pub,
		callerID:  callerID,
		tcpClient: tcpClient,
		udpAddr:   udpAddr,
		ctx:       ctx,
		ctxCancel: ctxCancel,
	}

	pub.subscribers[callerID] = ps

	pub.subscribersWg.Add(1)
	go ps.run()
}

func (ps *publisherSubscriber) close() {
	delete(ps.pub.subscribers, ps.callerID)
	ps.ctxCancel()
}

func (ps *publisherSubscriber) run() {
	defer ps.pub.subscribersWg.Done()

	if ps.tcpClient != nil {
		ps.runTCP()
	} else {
		ps.runUDP()
	}
}

func (ps *publisherSubscriber) runTCP() {
	if ps.pub.conf.onSubscriber != nil {
		ps.pub.conf.onSubscriber()
	}

	readerDone := make(chan struct{})
	go func() {
		defer close(readerDone)

		buf := make([]byte, 64)
		for {
			_, err := ps.tcpClient.NetConn().Read(buf)
			if err != nil {
				return
			}
		}
	}()

	select {
	case <-readerDone:
		ps.tcpClient.Close()

		select {
		case ps.pub.subscriberTCPClose <- ps:
		case <-ps.pub.ctx.Done():
		}

		<-ps.ctx.Done()

	case <-ps.ctx.Done():
		ps.tcpClient.Close()
		<-readerDone
	}
}

func (ps *publisherSubscriber) runUDP() {
	if ps.pub.conf.onSubscriber != nil {
		ps.pub.conf.onSubscriber()
	}

	<-ps.ctx.Done()
}

func (ps *publisherSubscriber) writeMessage(msg interface{}) {
	if ps.tcpClient != nil {
		ps.tcpClient.WriteMessage(msg)
	} else {
		ps.curMessageID++

		var rawMessage bytes.Buffer
		err := protocommon.MessageEncode(&rawMessage, msg)
		if err != nil {
			return
		}
		byts := rawMessage.Bytes()

		frames := protoudp.FramesForPayload(
			uint32(ps.pub.id),
			ps.curMessageID,
			byts)

		for _, f := range frames {
			ps.pub.conf.Node.udprosServer.WriteFrame(f, ps.udpAddr)
		}
	}
}
