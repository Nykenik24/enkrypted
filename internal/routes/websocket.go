package routes

import (
	"log/slog"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

func registerWebSocketRoutes(app *fiber.App) {
	app.Use("/ws", func(c fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		slog.Info("new ws connection")
		slog.Info("allowed", "value", c.Locals("allowed"))
		slog.Info("id", "value", c.Params("id"))
		slog.Info("query", "value", c.Query("v"))
		slog.Info("session", "value", c.Cookies("session"))

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				slog.Error("error reading message from client", "error", err)
				break
			}
			slog.Info("got message from ws client", "recv", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				slog.Error("error writing message to client", "error", err)
				break
			}
		}

	}))
}
