package goroslib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
	"strconv"

	"github.com/aler9/goroslib/apimaster"
	"github.com/aler9/goroslib/apislave"
	"github.com/aler9/goroslib/msgutils"
	"github.com/aler9/goroslib/protocommon"
	"github.com/aler9/goroslib/prototcp"
	"github.com/aler9/goroslib/protoudp"
)

type publisherRequestTopicReq struct {
	req *api_slave.RequestRequestTopic
	res chan api_slave.ResponseRequestTopic
}

type publisherSubscriberTcpNewReq struct {
	client *proto_tcp.Conn
	header *proto_tcp.HeaderSubscriber
}

// PublisherConf is the configuration of a Publisher.
type PublisherConf struct {
	// node which the publisher belongs to
	Node *Node

	// name of the topic in which messages will be written
	Topic string

	// an instance of the message that will be published
	Msg interface{}

	// (optional) whether to enable latching, that consists in saving the last
	// published message and send it to any new subscriber that connects to
	// this publisher
	Latch bool
}

// Publisher is a ROS publisher, an entity that can publish messages in a named channel.
type Publisher struct {
	conf        PublisherConf
	msgType     string
	msgMd5      string
	subscribers map[string]*publisherSubscriber
	lastMessage interface{}
	id          int

	requestTopic       chan publisherRequestTopicReq
	subscriberTcpNew   chan publisherSubscriberTcpNewReq
	subscriberTcpClose chan *publisherSubscriber
	write              chan interface{}
	shutdown           chan struct{}
	nodeTerminate      chan struct{}
	terminate          chan struct{}
	done               chan struct{}
}

// NewPublisher allocates a Publisher. See PublisherConf for the options.
func NewPublisher(conf PublisherConf) (*Publisher, error) {
	if conf.Node == nil {
		return nil, fmt.Errorf("Node is empty")
	}

	if len(conf.Topic) < 1 || conf.Topic[0] != '/' {
		return nil, fmt.Errorf("Topic must begin with /")
	}

	msgt := reflect.TypeOf(conf.Msg)
	if msgt.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("Msg must be a pointer")
	}
	if msgt.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("Msg must be a pointer to a struct")
	}

	msgType, err := msg_utils.Type(conf.Msg)
	if err != nil {
		return nil, err
	}

	msgMd5, err := msg_utils.Md5Message(conf.Msg)
	if err != nil {
		return nil, err
	}

	p := &Publisher{
		conf:               conf,
		msgType:            msgType,
		msgMd5:             msgMd5,
		subscribers:        make(map[string]*publisherSubscriber),
		requestTopic:       make(chan publisherRequestTopicReq),
		subscriberTcpNew:   make(chan publisherSubscriberTcpNewReq),
		subscriberTcpClose: make(chan *publisherSubscriber),
		write:              make(chan interface{}),
		shutdown:           make(chan struct{}),
		nodeTerminate:      make(chan struct{}),
		terminate:          make(chan struct{}),
		done:               make(chan struct{}),
	}

	chanErr := make(chan error)
	conf.Node.publisherNew <- publisherNewReq{
		pub: p,
		err: chanErr,
	}
	err = <-chanErr
	if err != nil {
		return nil, err
	}

	go p.run()

	return p, nil
}

// Close closes a Publisher and shuts down all its operations.
func (p *Publisher) Close() error {
	p.shutdown <- struct{}{}
	close(p.terminate)
	<-p.done
	return nil
}

func (p *Publisher) run() {
outer:
	for {
		select {
		case req := <-p.requestTopic:
			err := func() error {
				if len(req.req.Protocols) < 1 {
					return fmt.Errorf("invalid protocol")
				}

				proto := req.req.Protocols[0]

				if len(proto) < 1 {
					return fmt.Errorf("invalid protocol")
				}

				protoName, ok := proto[0].(string)
				if !ok {
					return fmt.Errorf("invalid protocol")
				}

				switch protoName {
				case "TCPROS":
					req.res <- api_slave.ResponseRequestTopic{
						Code:          1,
						StatusMessage: "",
						Protocol: []interface{}{
							"TCPROS",
							p.conf.Node.nodeIp.String(),
							p.conf.Node.tcprosServerPort,
						},
					}
					return nil

				case "UDPROS":
					if len(proto) < 5 {
						return fmt.Errorf("invalid protocol")
					}

					protoDef, ok := proto[1].([]byte)
					if !ok {
						return fmt.Errorf("invalid protoDef")
					}

					protoHost, ok := proto[2].(string)
					if !ok {
						return fmt.Errorf("invalid protoHost")
					}

					protoPort, ok := proto[3].(int)
					if !ok {
						return fmt.Errorf("invalid protoPort")
					}

					_, ok = proto[4].(int)
					if !ok {
						return fmt.Errorf("invalid proto1500")
					}

					newProtoDef := make([]byte, 4)
					binary.LittleEndian.PutUint32(newProtoDef, uint32(len(protoDef)))
					newProtoDef = append(newProtoDef, protoDef...)
					buf := bytes.NewBuffer(newProtoDef)

					raw, err := proto_common.HeaderDecodeRaw(buf)
					if err != nil {
						return err
					}

					var header proto_udp.HeaderSubscriber
					err = proto_common.HeaderDecode(raw, &header)
					if err != nil {
						return err
					}

					_, ok = p.subscribers[header.Callerid]
					if ok {
						return fmt.Errorf("topic '%s' is already subscribed by '%s'",
							p.conf.Topic, header.Callerid)
					}

					if header.Md5sum != p.msgMd5 {
						return fmt.Errorf("wrong md5: expected '%s', got '%s'",
							p.msgMd5, header.Md5sum)
					}

					udpAddr, err := net.ResolveUDPAddr("udp4", protoHost+":"+strconv.FormatInt(int64(protoPort), 10))
					if err != nil {
						return fmt.Errorf("unable to solve udp address)")
					}

					// if subscriber is in localhost, send packets from localhost to localhost
					// this avoids a bug in which the source ip is randomly chosen
					// from all available interfaces, making ip-based filtering unpractical
					isLocalhost := func() bool {
						ifaces, err := net.Interfaces()
						if err != nil {
							return false
						}

						for _, i := range ifaces {
							addrs, err := i.Addrs()
							if err != nil {
								continue
							}

							for _, addr := range addrs {
								if v, ok := addr.(*net.IPNet); ok {
									if v.IP.Equal(udpAddr.IP) {
										return true
									}
								}
							}
						}
						return false
					}()

					if isLocalhost {
						udpAddr.IP = net.IPv4(127, 0, 0, 1)
					}

					p.subscribers[header.Callerid] = newPublisherSubscriber(p,
						header.Callerid, nil, udpAddr)

					req.res <- api_slave.ResponseRequestTopic{
						Code:          1,
						StatusMessage: "",
						Protocol: []interface{}{
							"UDPROS",
							func() string {
								if isLocalhost {
									return "127.0.0.1"
								}
								return p.conf.Node.nodeIp.String()
							}(),
							p.conf.Node.udprosServerPort,
							p.id,
							1500,
							func() []byte {
								buf := bytes.NewBuffer(nil)
								proto_common.HeaderEncode(buf, &proto_udp.HeaderPublisher{
									Callerid: p.conf.Node.conf.Name,
									Md5sum:   p.msgMd5,
									Topic:    p.conf.Topic,
									Type:     p.msgType,
								})
								return buf.Bytes()[4:]
							}(),
						},
					}

					return nil
				}

				return fmt.Errorf("invalid protocol")
			}()
			if err != nil {
				req.res <- api_slave.ResponseRequestTopic{
					Code:          0,
					StatusMessage: err.Error(),
				}
				continue
			}

		case req := <-p.subscriberTcpNew:
			err := func() error {
				_, ok := p.subscribers[req.header.Callerid]
				if ok {
					return fmt.Errorf("topic '%s' is already subscribed by '%s'",
						p.conf.Topic, req.header.Callerid)
				}

				// wildcard is used by rostopic hz
				if req.header.Md5sum != "*" && req.header.Md5sum != p.msgMd5 {
					return fmt.Errorf("wrong md5: expected '%s', got '%s'",
						p.msgMd5, req.header.Md5sum)
				}

				err := req.client.WriteHeader(&proto_tcp.HeaderPublisher{
					Callerid: p.conf.Node.conf.Name,
					Md5sum:   p.msgMd5,
					Topic:    p.conf.Topic,
					Type:     p.msgType,
					Latching: func() int {
						if p.conf.Latch {
							return 1
						}
						return 0
					}(),
				})
				if err != nil {
					req.client.Close()
					return nil
				}

				p.subscribers[req.header.Callerid] = newPublisherSubscriber(p,
					req.header.Callerid, req.client, nil)

				if p.conf.Latch && p.lastMessage != nil {
					p.subscribers[req.header.Callerid].writeMessage(p.lastMessage)
				}

				return nil
			}()
			if err != nil {
				req.client.WriteHeader(&proto_tcp.HeaderError{
					Error: err.Error(),
				})
				req.client.Close()
				continue
			}

		case sub := <-p.subscriberTcpClose:
			delete(p.subscribers, sub.callerid)
			close(sub.terminate)

		case msg := <-p.write:
			if p.conf.Latch {
				p.lastMessage = msg
			}

			for _, s := range p.subscribers {
				s.writeMessage(msg)
			}

		case <-p.shutdown:
			break outer
		}
	}

	go func() {
		for {
			select {
			case req, ok := <-p.requestTopic:
				if !ok {
					return
				}

				req.res <- api_slave.ResponseRequestTopic{
					Code:          0,
					StatusMessage: "terminating",
				}

			case <-p.subscriberTcpNew:
			case <-p.subscriberTcpClose:
			case <-p.write:
			case <-p.shutdown:
			}
		}
	}()

	p.conf.Node.apiMasterClient.UnregisterPublisher(api_master.RequestUnregister{
		Topic:     p.conf.Topic[1:],
		CallerUrl: p.conf.Node.apiSlaveServerUrl,
	})

	for _, ps := range p.subscribers {
		close(ps.terminate)
		<-ps.done
	}

	p.conf.Node.publisherClose <- p

	<-p.nodeTerminate
	<-p.terminate

	close(p.requestTopic)
	close(p.subscriberTcpNew)
	close(p.subscriberTcpClose)
	close(p.write)
	close(p.shutdown)
	close(p.done)
}

// Write writes a message into the publisher.
func (p *Publisher) Write(msg interface{}) {
	if reflect.TypeOf(msg) != reflect.TypeOf(p.conf.Msg) {
		panic("wrong message type")
	}

	p.write <- msg
}
