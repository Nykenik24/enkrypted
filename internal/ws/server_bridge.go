package ws

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/services/auth"
)

type Server interface {
	GetRoom(id *id.ID) (Room, error)
	GetAllRooms() map[*id.ID]Room
	GetAuth() *auth.AuthService
	AddRoom() (Room, error)
	RemoveRoom(id *id.ID) error
}

type Room interface {
	Broadcast(ev *event.Event)
	Join(c *Client) error
	Leave(c *Client)
}
