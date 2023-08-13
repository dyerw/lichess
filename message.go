package main

import (
	"encoding/json"
	"errors"
)

/*
We should probably distinguish between server->client and client->server message types.
i.e. LobbyCreated is not a valid message client->server
*/
type Message interface {
	isMessage()
}

type CreateLobby struct{}

func (cl CreateLobby) isMessage() {}

type LobbyCreated struct {
	id string
}

func (lc LobbyCreated) isMessage() {}

type JoinLobby struct {
	id string
}

func (jl JoinLobby) isMessage() {}

type WebsocketMessage struct {
	MessageType string
	Payload     json.RawMessage
}

func UnmarshalMessage(b []byte) (Message, error) {
	var websocketMessage WebsocketMessage
	err := json.Unmarshal(b, &websocketMessage)
	if err != nil {
		return nil, err
	}

	switch websocketMessage.MessageType {
	case "CreateLobby":
		var createLobby CreateLobby
		json.Unmarshal(websocketMessage.Payload, &createLobby)
		return createLobby, nil
	case "JoinLobby":
		var joinLobby JoinLobby
		json.Unmarshal(websocketMessage.Payload, &joinLobby)
		return joinLobby, nil
	default:
		return nil, errors.New("Invalid MessageType")
	}
}
