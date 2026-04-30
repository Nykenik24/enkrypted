package app

import (
	"log"

	"github.com/Nykenik24/enkrypted/internal/db"
	"github.com/Nykenik24/enkrypted/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func Start() {
	app := fiber.New()

	db := db.GetInstance().Database
	database, err := db.DB()
	if err != nil {
		log.Fatalf("could not get db instance: %s", err.Error())
	}
	defer database.Close()

	routes.RegisterAll(app)

	log.Fatal(app.Listen(":8080"))
}
