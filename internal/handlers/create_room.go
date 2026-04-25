package handlers

import (
	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type CreateRoomHandler struct{}

func (h *CreateRoomHandler) Handle(ctx *ws.Context) error {
	var ev builtin_ev.CreateRoomEvent

	if err := ctx.Bind(&ev); err != nil {
		return err
	}

	ctx.Server.AddRoom()

	return nil
}
