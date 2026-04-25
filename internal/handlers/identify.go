package handlers

import (
	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type IdentifyHandler struct{}

func (h *IdentifyHandler) Handle(ctx *ws.Context) error {
	var ev builtin_ev.IdentifyEvent

	if err := ctx.BindData(&ev); err != nil {
		return err
	}

	if ev.Username == "" {
		ev.Username = user.GenericUsername()
	}

	ctx.Client.SetUsername(ev.Username)
	return nil
}
