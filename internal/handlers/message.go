package handlers

import (
	"time"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type MessageHandler struct{}

func (h *MessageHandler) Handle(ctx *ws.Context) error {
	var msg struct {
		RoomID   uint64 `json:"roomId"`
		Contents string `json:"contents"`
	}

	if err := ctx.BindData(&data); err != nil {
		return err
	}

	ev := builtin_ev.NewMessageEvent(
		msg.Contents,
		time.Now().Format(time.RFC3339),
		ctx.Client.GetUser(),
	)

	ctx.Broadcast(
		builtin_ev.Generic(ev).ToRoom(msg.RoomID),
	)

	return nil
}
