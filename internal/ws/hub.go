package ws

import (
	"fmt"
	"log"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/middleware"
)

type Hub struct {
	clients map[uint64]*Client

	register   chan *Client
	unregister chan *Client

	process   chan *Envelope
	broadcast chan *event.Event

	handlers map[string]EventHandler

	midwareManager *middleware.MiddlewareManager

	server Server
}

type Envelope struct {
	Client *Client
	Event  *event.Event
}

func NewHub(server Server) *Hub {
	return &Hub{
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		process:        make(chan *Envelope, 256),
		broadcast:      make(chan *event.Event, 256),
		clients:        make(map[uint64]*Client),
		handlers:       make(map[string]EventHandler),
		midwareManager: middleware.NewMidwareManager(),
		server:         server,
	}
}

func (h *Hub) SetServer(s Server) {
	h.server = s
}

func (h *Hub) Register(c *Client) {
	h.register <- c
}

func (h *Hub) Unregister(c *Client) {
	h.unregister <- c
}

func (h *Hub) dispatch(ev *event.Event) {
	if ev.Target == nil || ev.Target.Broadcast {
		for _, c := range h.clients {
			_ = c.SendEvent(ev)
		}
		return
	}

	if ev.Target.UserID != nil {
		if c, ok := h.clients[*ev.Target.UserID]; ok {
			_ = c.SendEvent(ev)
		}
		return
	}

	if ev.Target.RoomID != nil {
		room, err := h.server.GetRoom(*ev.Target.RoomID)
		if err != nil {
			log.Println(err)
			return
		}

		room.Broadcast(ev)
		return
	}
}
func (h *Hub) Run() {
	for {
		select {

		case c := <-h.register:
			h.clients[c.GetID()] = c

		case c := <-h.unregister:
			delete(h.clients, c.GetID())
			c.CloseSend()

		case env := <-h.process:
			h.handleEvent(env)

		case ev := <-h.broadcast:
			h.dispatch(ev)
		}
	}
}

func (h *Hub) handleEvent(env *Envelope) {
	ev := env.Event

	var err error
	ev, err = h.midwareManager.MultiInject(ev)
	if err != nil {
		log.Println("middleware error:", err)
		return
	}

	handler, ok := h.handlers[ev.Kind.String()]
	if !ok {
		log.Printf("no handler for %s", ev.Kind.String())
		return
	}

	ctx := &Context{
		Client: env.Client,
		Hub:    h,
		Event:  ev,
		Server: h.server,
	}

	if err := handler.Handle(ctx); err != nil {
		log.Printf("handler error: %v", err)
	}
}

func (h *Hub) Emit(c *Client, ev *event.Event) {
	h.process <- &Envelope{
		Client: c,
		Event:  ev,
	}
}

func (h *Hub) Broadcast(ev *event.Event) {
	h.broadcast <- ev
}

func (h *Hub) AddHandler(kind string, handler EventHandler) {
	h.handlers[kind] = handler
}

func (h *Hub) AddHandlers(handlers map[string]EventHandler) {
	for k, v := range handlers {
		h.handlers[k] = v
	}
}

func (h *Hub) BroadcastEvent(ev *event.Event) error {
	log.Printf("called BroadcastEvent")

	var err error = nil
	ev, err = h.midwareManager.MultiInject(ev)
	if err != nil {
		fmt.Printf("error when broadcasting event:\n--> \x1b[31m%v\x1b[0m", err)
		return err
	}

	rawJSON, err := ev.JSON()
	if err != nil {
		fmt.Printf("error when broadcasting event:\n--> \x1b[31m%v\x1b[0m", err)
		return err
	}

	log.Printf("sent event to broadcast channel: %s", rawJSON)
	h.broadcast <- ev

	log.Printf("broadcasted: %s", rawJSON)

	return nil
}

func (h *Hub) SendToClient(ev *event.Event, id uint64) error {
	if _, exists := h.clients[id]; !exists {
		return fmt.Errorf("client %d not registered in hub", id)
	}

	h.clients[id].SendEvent(ev)
	return nil
}

func (h *Hub) GetClient(id uint64) (*Client, error) {
	if _, exists := h.clients[id]; !exists {
		return nil, fmt.Errorf("tried to get user %d, but it's not in the hub", id)
	}

	return h.clients[id], nil
}
