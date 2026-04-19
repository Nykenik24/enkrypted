package server

import (
	"encoding/json"

	"github.com/Nykenik24/enkrypted/internal/event"
)

type Client interface {
	Send([]byte)
	CloseSend()
}

type Server struct {
	clients map[Client]bool

	register   chan Client
	unregister chan Client
	broadcast  chan []byte
}

func NewServer() *Server {
	return &Server{
		clients:    make(map[Client]bool),
		register:   make(chan Client),
		unregister: make(chan Client),
		broadcast:  make(chan []byte),
	}
}

func (s *Server) Register(c Client) {
	s.register <- c
}

func (s *Server) Unregister(c Client) {
	s.unregister <- c
}

func (s *Server) Broadcast(data []byte) {
	s.broadcast <- data
}

func (s *Server) Run() {
	for {
		select {

		case c := <-s.register:
			s.clients[c] = true

		case c := <-s.unregister:
			if _, ok := s.clients[c]; ok {
				delete(s.clients, c)
				c.CloseSend()
			}

		case msg := <-s.broadcast:
			for c := range s.clients {
				c.Send(msg)
			}
		}
	}
}

// helper (optional)
func (s *Server) BroadcastEvent(kind event.EventKind, payload any) error {
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

	s.Broadcast(data)
	return nil
}
