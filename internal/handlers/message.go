package handlers

import (
	"encoding/json"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type MessagePayload struct {
	Contents  string     `json:"contents"`
	Timestamp string     `json:"timestamp"`
	User      *user.User `json:"user"`
}

func MessageEV(c *ws.Client, payload json.RawMessage) error {
	var msg struct {
		Contents  string `json:"contents"`
		Timestamp string `json:"timestamp"`
	}

	if err := json.Unmarshal(payload, &msg); err != nil {
		return err
	}

	broadcast := MessagePayload{
		Contents:  msg.Contents,
		Timestamp: msg.Timestamp,
		User:      c.GetUser(),
	}

	return c.GetServer().BroadcastEvent(event.MessageEvent, broadcast)
}
