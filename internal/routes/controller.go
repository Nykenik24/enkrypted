package routes

import "github.com/gofiber/fiber/v3"

func RegisterAll(app *fiber.App) {

	registerWebSocketRoutes(app)

	apiGroup := app.Group("/api/v1")
	registerRoomRoutes(apiGroup)
}
