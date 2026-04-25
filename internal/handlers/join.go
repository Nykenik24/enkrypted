package handlers

import (
	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type JoinRoomHandler struct{}

func (h *JoinRoomHandler) Handle(ctx *ws.Context) error {
	var ev builtin_ev.JoinRoomEvent

	if err := ctx.Bind(&ev); err != nil {
		return err
	}

	room, err := ctx.Server.GetRoom(ev.RoomID)
	if err != nil {
		return err
	}

	if err := room.Join(ctx.Client); err != nil {
		return err
	}

	return nil
}
