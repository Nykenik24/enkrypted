package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/Nykenik24/enkrypted/internal/server"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

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
