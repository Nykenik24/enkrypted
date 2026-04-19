package handlers

import (
	"encoding/json"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

func MessageEV(c *ws.Client, payload json.RawMessage) error {
	var msg string
	if err := json.Unmarshal(payload, &msg); err != nil {
		return err
	}

	return c.GetServer().BroadcastEvent(event.MessageEvent, msg)
}
