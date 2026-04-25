package handlers

import (
	"fmt"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type CreateRoomHandler struct{}

func (h *CreateRoomHandler) Handle(ctx *ws.Context) error {
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

	_, err = ctx.Server.AddRoom()
	if err != nil {
		return err
	}

	builtin_ev.BroadcastMessage(ctx, "created room")

	return nil
}
