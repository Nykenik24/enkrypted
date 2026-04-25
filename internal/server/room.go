package server

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var lastRoomID uint64 = 0

type Status int

const (
	StatusConnected Status = iota
	StatusDisconnected
)

type Room struct {
	ID   uint64
	Host *Server

	members map[uint64]*ws.Client
}

func NewRoom(host *Server) *Room {
	lastRoomID++
	return &Room{
		Host: host,
		ID:   lastRoomID,

		members: make(map[uint64]*ws.Client),
	}
}

func (r *Room) Broadcast(ev *event.Event) {
	for _, c := range r.members {
		_ = c.SendEvent(ev)
	}
}

func (s *Server) AddRoom() (ws.Room, error) {
	room := NewRoom(s)
	s.Rooms[room.ID] = room

	return room, nil
}

func (s *Server) RemoveRoom(id uint64) error {
	if s.Rooms[id] == nil {
		return fmt.Errorf("room %d doesn't exist", id)
	}

	delete(s.Rooms, id)
	return nil
}
