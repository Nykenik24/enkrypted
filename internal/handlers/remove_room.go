package handlers

import (
	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type RemoveRoomHandler struct{}

func (h *RemoveRoomHandler) Handle(ctx *ws.Context) error {
	var ev builtin_ev.RemoveRoomEvent

	if err := ctx.Bind(&ev); err != nil {
		return err
	}

	ctx.Server.RemoveRoom(ev.RoomID)

	return nil
}
