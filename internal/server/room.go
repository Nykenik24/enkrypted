package server

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type Status int

const (
	StatusConnected Status = iota
	StatusDisconnected
)

type Room struct {
	ID   *id.ID  `json:"id"`
	Host *Server `json:"-"`

	members map[*id.ID]*ws.Client
}

func NewRoom(host *Server) *Room {
	return &Room{
		Host: host,
		ID:   id.RandomID(),

		members: make(map[*id.ID]*ws.Client),
	}
}

func (r *Room) Broadcast(ev *event.Event) {
	for _, c := range r.members {
		_ = c.SendEvent(ev)
	}
}

func (s *Server) AddRoom() (ws.Room, error) {
	room := NewRoom(s)
	s.Rooms[room.ID.String()] = room

	return room, nil
}

func (s *Server) RemoveRoom(id *id.ID) error {
	if s.Rooms[id.String()] == nil {
		return fmt.Errorf("room %s doesn't exist", id)
	}

	delete(s.Rooms, id.String())
	return nil
}
