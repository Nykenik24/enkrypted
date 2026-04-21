package ws

import (
	"fmt"
	"log"

	"github.com/Nykenik24/enkrypted/internal/event"
)

func (c *Client) AddHandler(kindStr string, handler EventHandler) {
	kind, err := event.KindFromString(kindStr)
	if err != nil {
		panic(err)
	}
	c.handlers[*kind] = handler
}

func (c *Client) RemoveHandler(kind event.EventKind) {
	delete(c.handlers, kind)
}

func (c *Client) Handle(ev *event.Event) error {
	rawJSON, err := ev.JSON()
	if err == nil {
		log.Printf("handling event: %s", rawJSON)
	}

	handler, ok := c.handlers[*ev.Kind]
	if !ok {
		return fmt.Errorf("handler for kind '%s' doesn't exist", ev.Kind.String())
	}

	err = handler.Handle(c, ev.Data)
	if err != nil {
		return fmt.Errorf("error when handling event:\n--> \x1b[31m%v\x1b[0m", err)
	}

	return nil
}
