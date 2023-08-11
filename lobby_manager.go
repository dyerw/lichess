package main

import "github.com/lesismal/nbio/nbhttp/websocket"

type LobbyManager struct {
}

func NewLobbyManager() *LobbyManager {
	return &LobbyManager{}
}

func (lm *LobbyManager) NewConnection(conn *websocket.Conn) {
	conn.OnMessage(func(c *websocket.Conn, mt websocket.MessageType, b []byte) {
		println("New message in lobby manager")
		message, err := ParseWebsocketMessage(b)
		if err != nil {
			println("PARSE ERROR")
		}
		println(message)
		switch v := message.(type) {
		case CreateLobby:
			println("CreateLobby")
		case JoinLobby:
			println("JoinLobby")
			println(v.id)
		}
	})
}
