package main

import "encoding/json"

type Message interface {
	isMessage()
	hasBody() bool
}

type CreateLobby struct{}

func (cl CreateLobby) isMessage()    {}
func (cl CreateLobby) hasBody() bool { return false }

type LobbyCreated struct {
	id string
}

func (lc LobbyCreated) isMessage()    {}
func (lc LobbyCreated) hasBody() bool { return true }

type JoinLobby struct {
	id string
}

func (jl JoinLobby) isMessage()    {}
func (jl JoinLobby) hasBody() bool { return true }

type WebsocketMessage struct {
	MessageType string
	Payload     json.RawMessage
}

func ParseWebsocketMessage(b []byte) (Message, error) {
	var websocketMessage WebsocketMessage
	err := json.Unmarshal(b, &websocketMessage)
	if err != nil {
		println("OHHHH NOOO")
	}

	var payload Message
	switch websocketMessage.MessageType {
	case "CreateLobby":
		payload = new(CreateLobby)
	case "JoinLobby":
		payload = new(JoinLobby)
	default:
		println("NOT MESSAGE")
	}
	if !payload.hasBody() {
		return payload, nil
	}
	payload_err := json.Unmarshal(websocketMessage.Payload, payload)
	if payload_err != nil {
		println("PAYLOAD ERROR")
	}
	return payload, nil
}
