package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/server"
	"github.com/Nykenik24/enkrypted/internal/user"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type EventHandler func(*Client, json.RawMessage) error

type Client struct {
	server *server.Server
	conn   *websocket.Conn
	send   chan []byte

	handlers map[event.EventKind]EventHandler
	user     *user.User
}

func NewClient(s *server.Server, conn *websocket.Conn) *Client {
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
	ev, err := event.NewAnyPayload(kind, payload)
	if err != nil {
		return err
	}

	c.Send(ev.JSON())
	return nil
}

func (c *Client) Handle(ev *event.Event) error {
	handler, ok := c.handlers[ev.Kind]
	if !ok {
		return fmt.Errorf("handler for kind '%d' doesn't exist", ev.Kind)
	}

	err := handler(c, ev.Payload)
	if err != nil {
		return fmt.Errorf("error when handling event:\n--> \x1b[31m%v\x1b[0m", err)
	}

	return nil
}

func ServeWebsockets(s *server.Server, w http.ResponseWriter, r *http.Request) *Client {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	c := NewClient(s, conn)

	s.Register(c)

	go c.ReadPump()
	go c.WritePump()

	return c
}
