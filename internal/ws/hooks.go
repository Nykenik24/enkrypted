package ws

import (
	"log/slog"
)

func RegisterDefaultEvents(c *Client) {
	c.OnConnect(func(ev *Event) (response *Event) {
		slog.Info("Client connected!")
		return ev
	})

	c.OnDisconnect(func(ev *Event) (response *Event) {
		slog.Info("Client disconnected!")
		return ev
	})
}
