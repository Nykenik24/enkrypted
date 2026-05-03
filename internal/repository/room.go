package repository

import (
	"context"
	"log/slog"

	"github.com/Nykenik24/enkrypted/internal/crypto"
	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/models"
	"gorm.io/gorm"
)

type RoomRepository = Repository[models.Room]

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
		slog.Error("error retrieving rooms", "error", err)
		return nil, err
	}

	slog.Info("retrieved all rows from rooms table", "count", result.RowsAffected)

	return rooms, nil
}

func (r *roomRepository) GetByID(id string) (*models.Room, error) {
	ctx := context.Background()

	room, err := gorm.G[models.Room](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		slog.Error("error getting room by id", "id", id, "error", err)
		return nil, err
	}

	slog.Info("retrieved row by id from rooms table", "id", id)

	return &room, nil
}

func (r *roomRepository) Create(room models.Room) (*models.Room, error) {
	ctx := context.Background()

	hashedPassword, err := crypto.HashPasswordSecure(room.Password)
	if err != nil {
		slog.Error("error hashing room password", "error", err)
		return nil, err
	}
	room.Password = hashedPassword

	err = gorm.G[models.Room](r.db).Create(ctx, &room)
	if err != nil {
		slog.Error("error inserting room into database", "room", room, "error", err)
		return nil, err
	}

	slog.Info("created new room", "room", room)

	return &room, nil
}

func (r *roomRepository) Delete(id string) (int, error) {
	ctx := context.Background()

	rowsAffected, err := gorm.G[models.Room](r.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		slog.Error("error deleting room from database", "id", id, "error", err)
		return -1, err
	}

	slog.Info("deleted room from database", "id", id)

	return rowsAffected, nil
}

func (r *roomRepository) Count() (int, error) {
	var rooms []models.Room
	result := r.db.Find(&rooms)

	if err := result.Error; err != nil {
		slog.Error("error retrieving rooms", "error", err)
		return -1, err
	}

	return len(rooms), nil
}

func (r *roomRepository) Clear() error {
	ctx := context.Background()

	_, err := gorm.G[models.Room](r.db).Delete(ctx)
	if err != nil {
		slog.Error("error clearing rooms from database", "error", err)
		return err
	}

	slog.Info("cleared rooms from database")

	return nil
}
