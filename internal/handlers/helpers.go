package handlers

import (
	"time"

	"github.com/Nykenik24/enkrypted/internal/ws"
)

func BroadcastMessage(ctx *ws.Context, contents string) {
	ctx.Broadcast(Base(
		NewMessageEvent(
			contents,
			time.Now().Format(time.RFC3339),
			ctx.Client.GetUser(),
		),
	))
}
