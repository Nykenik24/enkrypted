package server

import (
	"fmt"
	"log"

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
	Rooms map[string]*Room
	auth  *auth.AuthService
}

func NewServer(config *ServerConfig) *Server {
	authService, err := auth.NewAuthService(config.AdminPasswordHash)
	if err != nil {
		panic(err)
	}

	s := &Server{
		Rooms: make(map[string]*Room),
		auth:  authService,
	}

	h := ws.NewHub(s)
	h.SetServer(s)

	s.Hub = h

	return s
}

func (s *Server) GetAllRooms() map[*id.ID]ws.Room {
	rooms := make(map[*id.ID]ws.Room)

	for k, v := range s.Rooms {
		rooms[id.FromString(k)] = v
	}

	return rooms
}

func (s *Server) GetRoom(id *id.ID) (ws.Room, error) {
	log.Printf("\nrooms: %v\nroomId: %s", s.Rooms, id.String())

	room, ok := s.Rooms[id.String()]
	if !ok {
		return nil, fmt.Errorf("room %s not found", id)
	}

	return room, nil
}

func (s *Server) GetAuth() *auth.AuthService {
	return s.auth
}
