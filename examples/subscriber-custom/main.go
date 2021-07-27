package main

import (
	"fmt"

	"github.com/ROBTTO/goroslib"
	"github.com/ROBTTO/goroslib/pkg/msg"
)

// define a custom message.
// unlike the standard library, a .msg file is not needed.
// a structure definition is enough.
type TestMessage struct {
	msg.Package `ros:"my_package"`
	FirstField  uint32
	SecondField string
}

func onMessage(msg *TestMessage) {
	fmt.Printf("Incoming: %+v\n", msg)
}

func main() {
	// create a node and connect to the master
	n, err := goroslib.NewNode(goroslib.NodeConf{
		Name:          "goroslib_sub",
		MasterAddress: "127.0.0.1:11311",
	})
	if err != nil {
		panic(err)
	}
	defer n.Close()

	// create a subscriber
	sub, err := goroslib.NewSubscriber(goroslib.SubscriberConf{
		Node:     n,
		Topic:    "test_topic",
		Callback: onMessage,
	})
	if err != nil {
		panic(err)
	}
	defer sub.Close()

	// freeze main loop
	select {}
}
