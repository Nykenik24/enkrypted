package ws

import "github.com/Nykenik24/enkrypted/internal/event"

type Server interface {
	GetRoom(id uint64) (Room, error)
	AddRoom() (Room, error)
	RemoveRoom(id uint64) error
}

type Room interface {
	Broadcast(ev *event.Event)
	Join(c *Client) error
	Leave(c *Client)
}
