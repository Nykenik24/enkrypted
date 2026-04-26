package ws

import (
	"log"

	"github.com/gorilla/websocket"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/middleware"
	"github.com/Nykenik24/enkrypted/internal/user"
)

type EventHandler interface {
	Handle(*Context) (*event.Event, error)
	Kind() *event.EventKind
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte

	handlers map[string]EventHandler
	user     *user.User

	midwareManager *middleware.MiddlewareManager
}

func NewClient(s *Hub, conn *websocket.Conn) *Client {
	return &Client{
		hub:            s,
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
		c.Disconnect()
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

func (c *Client) Disconnect() {
	c.conn.Close()
	c.CloseSend()
	c.GetHub().Unregister(c)
}

func (c *Client) CloseSend() {
	close(c.send)
}

func (c *Client) GetHub() *Hub {
	return c.hub
}

func (c *Client) SetUsername(username string) {
	c.user.Username = username
}

func (c *Client) GetUser() *user.User {
	return c.user
}

// This function gets the ID from the user, meaning it is the equivalent of "client.GetUser().ID".
func (c *Client) GetID() *id.ID {
	return c.user.ID
}
