package db

import (
	"log"
	"sync"

	"github.com/Nykenik24/enkrypted/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type DB struct {
	Database *gorm.DB
}

var dbSingleton *DB

func CreateInstance() {
	lock.Lock()
	defer lock.Unlock()

	dbSingleton = &DB{Database: NewDB()}
}

func GetInstance() *DB {
	if dbSingleton == nil {
		CreateInstance()
	}
	return dbSingleton
}

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("enkrypted.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening db connection: %s", err.Error())
	}

	// ctx := context.Background()

	db.AutoMigrate(&models.Room{})

	return db
}
