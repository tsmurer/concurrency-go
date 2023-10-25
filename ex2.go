package main

import (
	"fmt"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch chan Message
}

func (s *Server) startAndListen() {
	for {
		select {

		case msg := <-s.msgch:
			fmt.Printf("msg received from: %s\npayload:\n%s", msg.From, msg.Payload)

		default:

		}

	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Tiago Murer",
		Payload: payload,
	}

	msgch <- msg
}

func messageSendingExample() {
	s := &Server{
		msgch: make(chan Message),
	}

	go s.startAndListen()

	sendMessageToServer(s.msgch, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

	select {}
}
