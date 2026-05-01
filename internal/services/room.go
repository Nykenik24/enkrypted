package services

import (
	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/Nykenik24/enkrypted/internal/repository"
)

type RoomService interface {
	GetAll() ([]models.Room, error)
	GetByID(id string) (*models.Room, error)
	Create(room models.Room) (*models.Room, error)
	Delete(id string) (int, error)
}

type roomService struct {
	repository repository.RoomRepository
}

func NewRoomService(roomRepository repository.RoomRepository) RoomService {
	return &roomService{repository: roomRepository}
}

func (s *roomService) GetAll() ([]models.Room, error) {
	return s.repository.GetAll()
}

func (s *roomService) GetByID(id string) (*models.Room, error) {
	return s.repository.GetByID(id)
}

func (s *roomService) Create(room models.Room) (*models.Room, error) {
	return s.repository.Create(room)
}

func (s *roomService) Delete(id string) (int, error) {
	return s.repository.Delete(id)
}
