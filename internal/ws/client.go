package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/middleware"
	"github.com/Nykenik24/enkrypted/internal/server"
	"github.com/Nykenik24/enkrypted/internal/user"
)

type EventHandler func(*Client, any) error

type Client struct {
	server *server.Server
	conn   *websocket.Conn
	send   chan []byte

	handlers map[event.EventKind]EventHandler
	user     *user.User

	midware map[string]middleware.Middleware
}

func NewClient(s *server.Server, conn *websocket.Conn) *Client {
	return &Client{
		server:   s,
		conn:     conn,
		send:     make(chan []byte, 256),
		handlers: make(map[event.EventKind]EventHandler),
		user:     user.NewUser(user.GenericUsername()),
		midware:  make(map[string]middleware.Middleware),
	}
}

func (c *Client) Send(data []byte) {
	select {
	case c.send <- data:
	default:
	}
}

func (c *Client) CloseSend() {
	close(c.send)
}

func (c *Client) GetServer() *server.Server {
	return c.server
}

func (c *Client) AddHandler(kind event.EventKind, handler EventHandler) {
	c.handlers[kind] = handler
}

func (c *Client) RemoveHandler(kind event.EventKind) {
	delete(c.handlers, kind)
}

func (c *Client) SetUsername(username string) {
	c.user.Username = username
}

func (c *Client) GetUser() *user.User {
	return c.user
}

func (c *Client) SendEvent(kind event.EventKind, payload any) error {
	ev := event.NewEvent(kind, payload)

	err := middleware.MultiInject(c.midware, ev)
	if err != nil {
		return err
	}

	rawJSON, err := ev.Marshal()
	if err != nil {
		return err
	}

	log.Printf("sending event: %s", rawJSON)
	c.Send(rawJSON)
	return nil
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

	err = handler(c, ev.Payload)
	if err != nil {
		return fmt.Errorf("error when handling event:\n--> \x1b[31m%v\x1b[0m", err)
	}

	return nil
}
