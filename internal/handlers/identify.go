package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type IdentifyPayload struct {
	Username string `json:"username"`
}

type IdentifyHandler struct{}

func (h *IdentifyHandler) Handle(c *ws.Client, payload *event.EventData) error {
	var p IdentifyPayload

	rawJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rawJSON, &p); err != nil {
		return err
	}

	if p.Username == "" {
		p.Username = user.GenericUsername()
	}

	c.SetUsername(p.Username)

	ev := builtin_ev.NewMessageEvent(
		fmt.Sprintf("user joined %s", p.Username),
		time.Now().Format(time.RFC3339),
		c.GetUser(),
	)
	return c.GetServer().BroadcastEvent(builtin_ev.ToGeneric(ev))
}
