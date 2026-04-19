package handlers

import (
	"encoding/json"

	"github.com/Nykenik24/enkrypted/internal/event"
	client "github.com/Nykenik24/enkrypted/internal/ws"
)

func PingEV(c *client.Client, payload json.RawMessage) error {
	return c.SendEvent(event.PingEvent, "Pong!")
}
