package handlers

import (
	"fmt"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type ConnectHandler struct{}

func (h *ConnectHandler) Handle(ctx *ws.Context) error {
	builtin_ev.BroadcastMessage(ctx, fmt.Sprintf("user joined %s", ctx.Client.GetUser().Username))

	return nil
}
