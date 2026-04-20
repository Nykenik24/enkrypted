package ws

import (
	"log"

	"github.com/gorilla/websocket"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/middleware"
	"github.com/Nykenik24/enkrypted/internal/server"
	"github.com/Nykenik24/enkrypted/internal/user"
)

type EventHandler interface {
	Handle(*Client, any) error
}

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

func (c *Client) CloseSend() {
	close(c.send)
}

func (c *Client) GetServer() *server.Server {
	return c.server
}

func (c *Client) SetUsername(username string) {
	c.user.Username = username
}

func (c *Client) GetUser() *user.User {
	return c.user
}
