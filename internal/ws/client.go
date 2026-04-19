package ws

import (
	"encoding/json"
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
	ev := struct {
		Kind    event.EventKind `json:"kind"`
		Payload any             `json:"payload"`
	}{
		Kind:    kind,
		Payload: payload,
	}

	data, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	c.Send(data)
	return nil
}

func readPump(c *Client) {
	defer func() {
		c.server.Unregister(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		var ev event.Event
		if err := json.Unmarshal(msg, &ev); err != nil {
			log.Println("bad event:", err)
			continue
		}

		handler, ok := c.handlers[ev.Kind]
		if !ok {
			log.Printf("handler for kind '%d' doesn't exist", ev.Kind)
			continue
		}

		err = handler(c, ev.Payload)
		if err != nil {
			log.Printf("error when handling event:\n--> \x1b[31m%v\x1b[0m", err)
		}
	}
}

func writePump(c *Client) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWebsockets(s *server.Server, w http.ResponseWriter, r *http.Request) *Client {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	c := NewClient(s, conn)

	s.Register(c)

	go readPump(c)
	go writePump(c)

	return c
}
