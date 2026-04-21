package ws

import (
	"log"

	"github.com/gorilla/websocket"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
)

type EventHandler interface {
	Handle(*Client, *event.EventData) error
}

type Client struct {
	server *Hub
	conn   *websocket.Conn
	send   chan []byte

	handlers map[event.EventKind]EventHandler
	user     *user.User
}

func NewClient(s *Hub, conn *websocket.Conn) *Client {
	return &Client{
		server:   s,
		conn:     conn,
		send:     make(chan []byte, 256),
		handlers: make(map[event.EventKind]EventHandler),
		user:     user.NewUser(user.GenericUsername()),
	}
}

func (c *Client) Send(data []byte) {
	select {
	case c.send <- data:
	default:
	}
}

func (c *Client) SendEvent(ev *event.Event) error {
	rawJSON, err := ev.JSON()
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

func (c *Client) GetServer() *Hub {
	return c.server
}

func (c *Client) SetUsername(username string) {
	c.user.Username = username
}

func (c *Client) GetUser() *user.User {
	return c.user
}
