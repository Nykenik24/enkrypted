package server

import (
	"github.com/Nykenik24/enkrypted/internal/ws"
)

func (r *Room) Join(c *ws.Client) error {
	r.members[c.GetID()] = c
	return nil
}

func (r *Room) Leave(c *ws.Client) {
	delete(r.members, c.GetID())
}
