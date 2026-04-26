package server

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/services/auth"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type ServerConfig struct {
	AdminPasswordHash string
}

func Config(adminPasswordHash string) *ServerConfig {
	return &ServerConfig{
		AdminPasswordHash: adminPasswordHash,
	}
}

type Server struct {
	Hub   *ws.Hub
	Rooms map[*id.ID]*Room
	auth  *auth.AuthService
}

func NewServer(config *ServerConfig) *Server {
	authService, err := auth.NewAuthService(config.AdminPasswordHash)
	if err != nil {
		panic(err)
	}

	s := &Server{
		Rooms: make(map[*id.ID]*Room),
		auth:  authService,
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

func (s *Server) GetAuth() *auth.AuthService {
	return s.auth
}
