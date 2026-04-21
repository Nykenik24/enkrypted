package ws

import (
	"fmt"
	"log"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/middleware"
)

func (c *Client) AddHandler(kind event.EventKind, handler EventHandler) {
	c.handlers[kind] = handler
}

func (c *Client) RemoveHandler(kind event.EventKind) {
	delete(c.handlers, kind)
}

func (c *Client) Handle(ev *event.Event) error {
	err := middleware.MultiInject(c.midware, ev)
	if err != nil {
		return err
	}

	rawJSON, err := ev.Marshal()
	if err == nil {
		log.Printf("handling event: %s", rawJSON)
	}

	handler, ok := c.handlers[ev.Kind]
	if !ok {
		return fmt.Errorf("handler for kind '%d' doesn't exist", ev.Kind)
	}

	err = handler.Handle(c, ev.Payload)
	if err != nil {
		return fmt.Errorf("error when handling event:\n--> \x1b[31m%v\x1b[0m", err)
	}

	return nil
}
