package main

import (
	"github.com/lesismal/nbio/nbhttp/websocket"
)

type LobbyManager struct {
}

func NewLobbyManager() *LobbyManager {
	return &LobbyManager{}
}

func (lm *LobbyManager) NewConnection(conn *websocket.Conn) {
	conn.OnMessage(func(c *websocket.Conn, mt websocket.MessageType, b []byte) {
		message, err := UnmarshalMessage(b)
		if err != nil {
			println(err.Error())
		}
		switch v := message.(type) {
		case CreateLobby:
			println("CreateLobby")
		case JoinLobby:
			println("JoinLobby")
			println(v.id)
		}
	})
}
