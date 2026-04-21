package ws

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/gorilla/websocket"
)

func (c *Client) ReadPump() {
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

		ev, err := event.EventFromJSON(msg)
		if err != nil {
			log.Println("bad event:", err)
			var buf bytes.Buffer
			if err := json.Indent(&buf, msg, "", "  "); err != nil {
				log.Println("error when indenting:", err)
			}
			log.Println("raw JSON:", buf.String())
			continue
		}

		c.Handle(ev)
	}
}

func (c *Client) WritePump() {
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
