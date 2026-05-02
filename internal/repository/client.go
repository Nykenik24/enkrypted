package repository

import (
	"slices"

	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type ClientRepo struct {
	clients map[string]*ws.Client
}

func NewClientRepo() *ClientRepo {
	return &ClientRepo{
		clients: make(map[string]*ws.Client),
	}
}

var clientRepo *ClientRepo

func GlobalClientRepo() *ClientRepo {
	if clientRepo == nil {
		clientRepo = NewClientRepo()
	}

	return clientRepo
}

func (cr *ClientRepo) GetAll() map[string]*ws.Client {
	return cr.clients
}

func (cr *ClientRepo) Keys() []string {
	var keys []string
	for k := range cr.clients {
		keys = append(keys, k)
	}
	return keys
}

func (cr *ClientRepo) GetByID(id *id.ID) *ws.Client {
	c, exists := cr.clients[id.String()]
	if !exists {
		return nil
	}
	return c
}

func (cr *ClientRepo) Add(c *ws.Client) {
	if !slices.Contains(cr.Keys(), c.ID.String()) {
		cr.clients[c.ID.String()] = c
	}
}

func (cr *ClientRepo) Remove(id *id.ID) {
	delete(cr.clients, id.String())
}

func (cr *ClientRepo) Count() int {
	count := 0
	for range cr.clients {
		count++
	}

	return count
}
