package ws

import (
	"errors"
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/gofiber/contrib/v3/websocket"
)

type Client struct {
	conn     *websocket.Conn
	handlers map[string]EventHandler
	ID       *models.ID
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:     conn,
		handlers: make(map[string]EventHandler),
		ID:       models.RandomID(),
	}
}

func (c *Client) ReadRaw() (mt int, msg []byte, err error) {
	if mt, msg, err = c.conn.ReadMessage(); err != nil {
		return websocket.CloseMessage, nil, err
	}

	if mt != websocket.TextMessage {
		return websocket.CloseMessage, nil, errors.New("message must be a TextMessage")
	}

	return mt, msg, nil
}

func (c *Client) ReadEvent() (mt int, ev *Event, err error) {
	mt, msg, err := c.ReadRaw()

	ev, err = UnmarshalEvent(msg)
	if err != nil {
		return websocket.CloseMessage, nil, err
	}

	return mt, ev, nil
}

func (c *Client) On(kind string, f EventHandler) error {
	if !LookupKind(kind) {
		return fmt.Errorf("unknown kind: '%s'", kind)
	}

	c.handlers[kind] = f
	return nil
}

func (c *Client) OnConnect(f EventHandler) {
	c.On(ConnectEvent, f)
}

func (c *Client) OnDisconnect(f EventHandler) {
	c.On(DisconnectEvent, func(ev *Event) (response *Event) {
		c.Disconnect()
		return f(ev)
	})
}

func (c *Client) Handle(ev *Event) (*Event, error) {
	handler, exists := c.handlers[ev.Kind]
	if !exists {
		return nil, fmt.Errorf("no handler for %s", ev.Kind)
	}

	return handler(ev), nil
}

func (c *Client) HandleRaw(msg []byte) (*Event, error) {
	ev, err := UnmarshalEvent(msg)
	if err != nil {
		return nil, err
	}

	return c.Handle(ev)
}

func (c *Client) Reply(mt int, ev *Event) error {
	rawJSON, err := ev.JSON()
	if err != nil {
		return err
	}

	if err := c.WriteRaw(mt, rawJSON); err != nil {
		return err
	}

	return nil
}

func (c *Client) WriteRaw(mt int, msg []byte) error {
	if err := c.conn.WriteMessage(mt, msg); err != nil {
		return err
	}

	return nil
}

func (c *Client) Disconnect() {
	c.conn.Close()
}
