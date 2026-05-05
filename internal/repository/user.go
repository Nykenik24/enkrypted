package repository

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/Nykenik24/enkrypted/internal/util"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByName(name string) (*models.User, error)
	Create(client models.User) (*models.User, error)
	DeleteByID(id string) (int, error)
	DeleteByName(name string) (int, error)
	Count() (int, error)
	Clear() error
}

type userRepository struct {
	users []models.User
}

func NewUserRepo() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	return r.users, nil
}

func (r *userRepository) GetByID(id string) (*models.User, error) {
	for _, u := range r.users {
		if u.ID.CompareString(id) {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("user %s not found; couldn't be retrieved", id)
}

func (r *userRepository) GetByName(name string) (*models.User, error) {
	for _, u := range r.users {
		if u.Username == name {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("user %s not found; couldn't be retrieved", name)
}

func (r *userRepository) Create(u models.User) (*models.User, error) {
	r.users = append(r.users, u)
	return &r.users[len(r.users)-1], nil
}

func (r *userRepository) DeleteByID(id string) (int, error) {
	for i, u := range r.users {
		if u.ID.CompareString(id) {
			newUsers, err := util.RemoveByIndex(r.users, i)
			if err != nil {
				return -1, err
			}

			r.users = newUsers
			return i, nil
		}
	}

	return -1, fmt.Errorf("client %s not found; couldn't be deleted", id)
}

func (r *userRepository) DeleteByName(name string) (int, error) {
	for i, u := range r.users {
		if u.Username == name {
			newUsers, err := util.RemoveByIndex(r.users, i)
			if err != nil {
				return -1, err
			}

			r.users = newUsers
			return i, nil
		}
	}

	return -1, fmt.Errorf("client %s not found; couldn't be deleted", name)
}

func (r *userRepository) Count() (int, error) {
	return len(r.users), nil
}

func (r *userRepository) Clear() error {
	r.users = []models.User{}
	return nil
}
