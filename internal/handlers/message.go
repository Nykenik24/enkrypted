package handlers

import (
	"encoding/json"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type MessageHandler struct{}

func (h *MessageHandler) Handle(c *ws.Client, payload *event.EventData) error {
	var msg struct {
		Contents  string `json:"contents"`
		Timestamp string `json:"timestamp"`
	}

	rawJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rawJSON, &msg); err != nil {
		return nil
	}

	ev := builtin_ev.NewMessageEvent(
		msg.Contents,
		msg.Timestamp,
		c.GetUser(),
	)

	return c.GetServer().BroadcastEvent(builtin_ev.ToGeneric(ev))
}
