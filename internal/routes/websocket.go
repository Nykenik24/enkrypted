package routes

import (
	"log/slog"

	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/repository"
	"github.com/Nykenik24/enkrypted/internal/ws"
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

func emitError(err error) *ws.Event {
	ev := ws.NewEvent(ws.ErrorEvent)
	ev.ID = id.RandomID()
	errstr := err.Error()
	ev.Error = &errstr
	return ev
}

var repo = repository.GlobalClientRepo()

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

		ct := ws.NewClient(c)
		ws.RegisterDefaultEvents(ct)

		repo.Create(*ct)

		defer (func() {
			ct.Disconnect()
		})()

		var (
			mt  int       // message type
			ev  *ws.Event // message
			err error     // error
		)
		for {
			if mt, ev, err = ct.ReadEvent(); err != nil {
				slog.Error("error reading message from client", "error", err)
				ct.Reply(mt, emitError(err))
				continue
			}
			slog.Info("got message from ws client", "recv", "\n"+ev.String())

			if ev.ID == nil {
				ev.ID = id.RandomID()
			}

			reply, err := ct.Handle(ev)
			if err != nil {
				slog.Error("error handling event", "error", err)
				ct.Reply(mt, emitError(err))
				continue
			}
			slog.Info("replying to client", "reply", "\n"+reply.String())

			if ev.Kind == ws.DisconnectEvent {
				repo.Delete(ct.ID.String())
				slog.Info("Client disconnected", "client", string(ct.ID.Short()))
				break
			}

			if err := ct.Reply(mt, reply); err != nil {
				slog.Error("error writing message to client", "error", err)
				ct.Reply(mt, emitError(err))
				continue
			}
		}

	}))
}
