package repository

import (
	"context"
	"log"
	"log/slog"

	"github.com/Nykenik24/enkrypted/internal/crypto"
	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/models"
	"gorm.io/gorm"
)

type RoomRepository interface {
	GetAll() ([]models.Room, error)
	GetByID(id string) (*models.Room, error)
	Create(room models.Room) (*models.Room, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository() RoomRepository {
	return &roomRepository{db: db.GetInstance().Database}
}

func (r *roomRepository) GetAll() ([]models.Room, error) {
	var rooms []models.Room
	result := r.db.Find(&rooms)

	if err := result.Error; err != nil {
		log.Printf("error retrieving rooms: %s", err.Error())
	}

	slog.Info("retrieved all rows from rooms table", "count", result.RowsAffected)

	return rooms, nil
}

func (r *roomRepository) GetByID(id string) (*models.Room, error) {
	ctx := context.Background()

	room, err := gorm.G[models.Room](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		log.Printf("error getting room with id: %s, error: %s", id, err.Error())
		return nil, err
	}

	slog.Info("retrieved row by id from rooms table", "id", id)

	return &room, nil
}

func (r *roomRepository) Create(room models.Room) (*models.Room, error) {
	ctx := context.Background()

	hashedPassword, err := crypto.HashPasswordSecure(room.Password)
	if err != nil {
		log.Printf("error hashing room password: %s", err.Error())
		return nil, err
	}
	room.Password = hashedPassword

	err = gorm.G[models.Room](r.db).Create(ctx, &room)
	if err != nil {
		log.Printf("error inserting room into database: %s", err.Error())
		return nil, err
	}

	slog.Info("created new room", "room", room)

	return &room, nil
}
