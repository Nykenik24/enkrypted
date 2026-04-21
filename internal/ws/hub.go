package ws

import (
	"log"
	"slices"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/middleware"
)

type Hub struct {
	clients []*Client

	register   chan Client
	unregister chan Client
	broadcast  chan *event.Event

	midwareManager *middleware.MiddlewareManager
}

func NewHub() *Hub {
	return &Hub{
		register:       make(chan Client),
		unregister:     make(chan Client),
		broadcast:      make(chan *event.Event),
		midwareManager: middleware.NewMidwareManager(),
	}
}

func (h *Hub) Register(c *Client) {
	h.register <- *c
}

func (h *Hub) Unregister(c *Client) {
	h.unregister <- *c
}

func (h *Hub) Run() {
	for {
		select {

		case c := <-h.register:
			h.clients = append(h.clients, &c)

		case c := <-h.unregister:
			for i, v := range h.clients {
				if v == &c {
					h.clients = slices.Delete(h.clients, i, i)
				}
			}

		case msg := <-h.broadcast:
			for _, c := range h.clients {
				c.SendEvent(msg)
			}
		}
	}
}

func (h *Hub) BroadcastEvent(ev *event.Event) error {
	var err error = nil
	ev, err = h.midwareManager.MultiInject(ev)
	if err != nil {
		return err
	}

	h.broadcast <- ev

	data, err := ev.JSON()
	if err != nil {
		return err
	}
	log.Printf("broadcasted: %s", data)

	return nil
}
