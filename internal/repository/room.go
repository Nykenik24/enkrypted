package repository

import (
	"context"
	"fmt"
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
	Delete(id string) (int, error)
	Count() (int, error)
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepo() RoomRepository {
	return &roomRepository{db: db.GetInstance().Database}
}

func (r *roomRepository) GetAll() ([]models.Room, error) {
	var rooms []models.Room
	result := r.db.Find(&rooms)

	if err := result.Error; err != nil {
		return nil, fmt.Errorf("error retrieving rooms: %v", err)
	}

	slog.Info("retrieved all rows from rooms table", "count", result.RowsAffected)

	return rooms, nil
}

func (r *roomRepository) GetByID(id string) (*models.Room, error) {
	ctx := context.Background()

	room, err := gorm.G[models.Room](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting room by id (%s): %v", id, err)
	}

	slog.Info("retrieved row by id from rooms table", "id", id)

	return &room, nil
}

func (r *roomRepository) Create(room models.Room) (*models.Room, error) {
	ctx := context.Background()

	hashedPassword, err := crypto.HashPasswordSecure(room.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing room password: %v", err)
	}
	room.Password = hashedPassword

	err = gorm.G[models.Room](r.db).Create(ctx, &room)
	if err != nil {
		return nil, fmt.Errorf("error inserting room into database: %v", err)
	}

	slog.Info("created new room", "room", room)

	return &room, nil
}

func (r *roomRepository) Delete(id string) (int, error) {
	ctx := context.Background()

	rowsAffected, err := gorm.G[models.Room](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		return -1, fmt.Errorf("error deleting room %s from database: %v", id, err)
	}

	slog.Info("deleted room from database", "id", id)

	return rowsAffected, nil
}

func (r *roomRepository) Count() (int, error) {
	var rooms []models.Room
	result := r.db.Find(&rooms)

	if err := result.Error; err != nil {
		return -1, fmt.Errorf("error retrieving rooms: %v", err)
	}

	return len(rooms), nil
}

func (r *roomRepository) Clear() error {
	ctx := context.Background()

	_, err := gorm.G[models.Room](r.db).Delete(ctx)
	if err != nil {
		return fmt.Errorf("error clearing rooms from database: %v", err)
	}

	slog.Info("cleared rooms from database")

	return nil
}
