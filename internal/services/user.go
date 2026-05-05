package services

import (
	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/Nykenik24/enkrypted/internal/repository"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	GetByName(name string) (*models.User, error)
	Create(User models.User) (*models.User, error)
	Delete(id string) (int, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{repository: userRepository}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repository.GetAll()
}

func (s *userService) GetByID(id string) (*models.User, error) {
	return s.repository.GetByID(id)
}

func (s *userService) GetByName(name string) (*models.User, error) {
	return s.repository.GetByName(name)
}

func (s *userService) Create(user models.User) (*models.User, error) {
	return s.repository.Create(user)
}

func (s *userService) Delete(id string) (int, error) {
	return s.repository.Delete(id)
}
