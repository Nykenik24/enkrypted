package routes

import (
	"github.com/Nykenik24/enkrypted/internal/handlers"
	"github.com/gofiber/fiber/v3"
)

func registerRoomRoutes(apiGroup fiber.Router) {
	roomGroup := apiGroup.Group("/rooms")

	roomGroup.Get("/all", handlers.GetAllRooms)
	roomGroup.Post("/new", handlers.CreateRoom)
}
