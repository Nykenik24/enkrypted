package handlers

import (
	"fmt"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type ConnectHandler struct{}

func (h *ConnectHandler) Handle(ctx *ws.Context) error {
	ev := builtin_ev.NewMessageEvent(
		fmt.Sprintf("user joined %s", ctx.Client.GetUser().Username),
		time.Now().Format(time.RFC3339),
		ctx.Client.GetUser(),
	)

	ctx.Broadcast(
		builtin_ev.Generic(ev).ToAll(),
	)

	return nil
}
