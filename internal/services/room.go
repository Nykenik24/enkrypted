package services

import (
	"context"
	"log"

	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/models"
	"gorm.io/gorm"
)

func GetAllRooms() ([]models.Room, error) {
	db := db.GetInstance().Database

	var rooms []models.Room
	result := db.Find(&rooms)

	if err := result.Error; err != nil {
		log.Printf("error retrieving rooms: %s", err.Error())
	}

	log.Printf("retrieved %d rows from rooms table", result.RowsAffected)

	return rooms, nil
}

func CreateRoom(room models.Room) error {
	db := db.GetInstance().Database
	ctx := context.Background()

	err := gorm.G[models.Room](db).Create(ctx, &room)
	if err != nil {
		log.Printf("error inserting room into database: %s", err.Error())
		return err
	}

	return nil
}
