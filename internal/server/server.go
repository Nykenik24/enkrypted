package server

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/services/auth"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type Server struct {
	Hub   *ws.Hub
	Rooms map[uint64]*Room
	Auth  *auth.AuthService
}

func NewServer() *Server {
	s := &Server{
		Rooms: make(map[uint64]*Room),
		Auth:  auth.NewAuthService(),
	}

	h := ws.NewHub(s)
	h.SetServer(s)

	s.Hub = h

	return s
}

func (s *Server) GetRoom(id uint64) (ws.Room, error) {
	room, ok := s.Rooms[id]
	if !ok {
		return nil, fmt.Errorf("room %d not found", id)
	}
	return room, nil
}
