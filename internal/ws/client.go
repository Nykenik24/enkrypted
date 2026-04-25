package ws

import (
	"log"

	"github.com/gorilla/websocket"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/middleware"
	"github.com/Nykenik24/enkrypted/internal/user"
)

type EventHandler interface {
	Handle(*Context) error
}

type Client struct {
	server *Hub
	conn   *websocket.Conn
	send   chan []byte

	handlers map[string]EventHandler
	user     *user.User

	midwareManager *middleware.MiddlewareManager
}

func NewClient(s *Hub, conn *websocket.Conn) *Client {
	return &Client{
		server:         s,
		conn:           conn,
		send:           make(chan []byte, 256),
		handlers:       make(map[string]EventHandler),
		user:           user.NewUser(user.GenericUsername()),
		midwareManager: middleware.NewMidwareManager(),
	}
}

func (c *Client) Send(data []byte) {
	select {
	case c.send <- data:
	default:
		c.CloseSend()
		c.conn.Close()
	}
}

func (c *Client) SendEvent(ev *event.Event) error {
	var err error = nil
	ev, err = c.midwareManager.MultiInject(ev)
	if err != nil {
		return err
	}

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

func (c *Client) GetHub() *Hub {
	return c.server
}

func (c *Client) SetUsername(username string) {
	c.user.Username = username
}

func (c *Client) GetUser() *user.User {
	return c.user
}

// This function gets the ID from the user, meaning it is the equivalent of "client.GetUser().ID".
func (c *Client) GetID() uint64 {
	return c.user.ID
}
