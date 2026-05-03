package repository

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/util"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type ClientRepository interface {
	GetAll() ([]ws.Client, error)
	GetByID(id string) (*ws.Client, error)
	Create(client ws.Client) (*ws.Client, error)
	Delete(id string) (int, error)
	Count() (int, error)
}

type clientRepository struct {
	clients []ws.Client
}

func NewClientRepo() ClientRepository {
	return &clientRepository{}
}

func (r *clientRepository) GetAll() ([]ws.Client, error) {
	return r.clients, nil
}

func (r *clientRepository) GetByID(id string) (*ws.Client, error) {
	for _, c := range r.clients {
		if c.ID.CompareString(id) {
			return &c, nil
		}
	}

	return nil, fmt.Errorf("client %s not found; couldn't be retrieved", id)
}

func (r *clientRepository) Create(c ws.Client) (*ws.Client, error) {
	r.clients = append(r.clients, c)
	return &r.clients[len(r.clients)-1], nil
}

func (r *clientRepository) Delete(id string) (int, error) {
	for i, c := range r.clients {
		if c.ID.CompareString(id) {
			clients, err := util.RemoveByIndex(r.clients, i)
			if err != nil {
				return -1, err
			}

			r.clients = clients
			return i, nil
		}
	}

	return -1, fmt.Errorf("client %s not found; couldn't be deleted", id)
}

func (r *clientRepository) Count() (int, error) {
	return len(r.clients), nil
}

func (r *clientRepository) Clear() error {
	r.clients = []ws.Client{}
	return nil
}

var clientRepo ClientRepository

func GlobalClientRepo() ClientRepository {
	if clientRepo == nil {
		clientRepo = NewClientRepo()
	}

	return clientRepo
}
