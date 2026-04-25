package handlers

import (
	"fmt"
	"time"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type RemoveRoomHandler struct{}

func (h *RemoveRoomHandler) Handle(ctx *ws.Context) error {
	ev := ctx.Event

	if !ev.HasAuth() {
		return fmt.Errorf("%s must have auth information", builtin_ev.CreateRoomEventKind.String())
	}

	passwd, err := ev.GetPassword()
	if err != nil {
		return err
	}

	auth := ctx.Server.GetAuth()
	if !auth.Hasher.VerifyPassword(passwd, auth.GetAdminHash()) {
		return fmt.Errorf("wrong password")
	}

	var data builtin_ev.RemoveRoomEvent
	if err := ctx.BindData(&data); err != nil {
		return err
	}

	err = ctx.Server.RemoveRoom(data.RoomID)
	if err != nil {
		return err
	}

	ctx.Broadcast(builtin_ev.Generic(
		builtin_ev.NewMessageEvent(
			"removed room",
			time.Now().Format(time.RFC3339),
			ctx.Client.GetUser(),
		),
	))

	return nil
}
