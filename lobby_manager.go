package main

import (
	"math/rand"
	"sync"

	"github.com/lesismal/nbio/nbhttp/websocket"
)

// TODO: Probably move this to some util file, ripped from SO
var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type LobbyManager struct {
	lobbies    map[string]*Lobby
	lobbyiesMx sync.Mutex
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
			lm.CreateLobby(c)
		case JoinLobby:
			println("JoinLobby")
			println(v.id)
			// TODO: Implement LobbyJoin
		}
	})
}

func (lm *LobbyManager) CreateLobby(c *websocket.Conn) {
	lobby := NewLobby(c)

	lm.lobbyiesMx.Lock()
	defer lm.lobbyiesMx.Unlock()

	lobbyCode := randSeq(5)
	lm.lobbies[lobbyCode] = lobby

	// TODO: Send Lobby code to client/confirm lobby creation
}
