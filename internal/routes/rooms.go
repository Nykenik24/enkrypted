package routes

import (
	"github.com/Nykenik24/enkrypted/internal/handlers"
	"github.com/Nykenik24/enkrypted/internal/repository"
	"github.com/Nykenik24/enkrypted/internal/services"
	"github.com/gofiber/fiber/v3"
)

func registerRoomRoutes(apiGroup fiber.Router) {
	roomRepository := repository.NewRoomRepository()
	roomService := services.NewRoomService(roomRepository)
	roomHandler := handlers.NewRoomHandler(roomService)

	roomGroup := apiGroup.Group("/rooms")

	roomGroup.Get("", roomHandler.GetAll)
	roomGroup.Get(":id", roomHandler.GetByID)
	roomGroup.Post("", roomHandler.Create)
}
