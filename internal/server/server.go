package server

import (
	"encoding/json"
	"log"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/middleware"
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

	midware map[string]middleware.Middleware
}

func NewServer() *Server {
	return &Server{
		clients:    make(map[Client]bool),
		register:   make(chan Client),
		unregister: make(chan Client),
		broadcast:  make(chan []byte),
		midware:    make(map[string]middleware.Middleware),
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

func (s *Server) BroadcastEvent(kind event.EventKind, payload any) error {
	ev := event.NewEvent(kind, payload)

	err := middleware.MultiInject(s.midware, ev)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(ev, "", "\t")
	if err != nil {
		return err
	}

	log.Printf("broadcasted: %s", data)

	s.Broadcast(data)
	return nil
}
