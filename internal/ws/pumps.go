package ws

import (
	"log"
	"time"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/gorilla/websocket"
)

func (c *Client) ReadPump() {
	defer func() {
		c.hub.Unregister(c)
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		ev, err := event.EventFromJSON(msg)
		if err != nil {
			log.Println("bad event:", err)
			continue
		}

		if ev.ID == nil {
			ev.ID = id.RandomID()
		}

		log.Printf("event id=%s", ev.ID.String())
		c.hub.Emit(c, ev)
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
