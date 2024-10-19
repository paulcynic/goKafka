package main

import "github.com/gorilla/websocket"

func foo() {
	websocket.DefaultDialer.Dial("ws://foo", nil)
}

type Consumer interface {
	Start() error
}
