package routes

import (
	"github.com/Nykenik24/enkrypted/internal/handlers"
	"github.com/Nykenik24/enkrypted/internal/repository"
	"github.com/Nykenik24/enkrypted/internal/services"
	"github.com/gofiber/fiber/v3"
)

func registerUserRoutes(apiGroup fiber.Router) {
	userRepository := repository.NewUserRepo()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	userGroup := apiGroup.Group("/users")

	userGroup.Get("", userHandler.GetAll)
	userGroup.Get(":id", userHandler.GetByID)
	userGroup.Get("/@:name", userHandler.GetByName)
	userGroup.Post("", userHandler.Create)
	userGroup.Delete(":id", userHandler.Delete)
}
