package handlers

import (
	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type LeaveRoomHandler struct{}

func (h *LeaveRoomHandler) Handle(ctx *ws.Context) error {
	var ev builtin_ev.LeaveRoomEvent

	if err := ctx.Bind(&ev); err != nil {
		return err
	}

	room, err := ctx.Server.GetRoom(ev.RoomID)
	if err != nil {
		return err
	}

	room.Leave(ctx.Client)

	return nil
}
