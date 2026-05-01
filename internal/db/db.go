package db

import (
	"log/slog"
	"sync"

	"github.com/Nykenik24/enkrypted/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbSingleton *DB
	once        sync.Once
)

type DB struct {
	Database *gorm.DB
}

func createInstance() {
	once.Do(func() {
		dbSingleton = &DB{Database: NewDB()}
	})
}

func GetInstance() *DB {
	if dbSingleton == nil {
		createInstance()
	}
	return dbSingleton
}

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("enkrypted.db"), &gorm.Config{})
	if err != nil {
		slog.Error("error opening database connection", "error", err)
	}

	// ctx := context.Background()

	db.AutoMigrate(&models.Room{})

	return db
}
