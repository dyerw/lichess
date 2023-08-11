package main

import "github.com/lesismal/nbio/nbhttp/websocket"

type Lobby struct {
	firstPlayerConn  *websocket.Conn
	secondPlayerConn *websocket.Conn
}

func NewLobby(fpc *websocket.Conn) *Lobby {
	return &Lobby{firstPlayerConn: fpc}
}
