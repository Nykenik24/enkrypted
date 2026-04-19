package handlers

import (
	"encoding/json"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type IdentifyPayload struct {
	Username string `json:"username"`
}

func IdentifyEV(c *ws.Client, payload json.RawMessage) error {
	var p IdentifyPayload

	if err := json.Unmarshal(payload, &p); err != nil {
		return err
	}

	if p.Username == "" {
		p.Username = user.GenericUsername()
	}

	c.SetUsername(p.Username)

	return c.GetServer().BroadcastEvent(event.MessageEvent, map[string]any{
		"system":  true,
		"message": "user joined: " + p.Username,
	})
}
