package handlers

import (
	"fmt"
	"time"

	builtin_ev "github.com/Nykenik24/enkrypted/internal/builtin/events"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type MessageHandler struct{}

func (h *MessageHandler) Handle(ctx *ws.Context) error {
	var data builtin_ev.MessageEvent

	if err := ctx.BindData(&data); err != nil {
		return err
	}

	err := builtin_ev.ValidateMessage(&data)
	if err != nil {
		return err
	}

	if data.RoomID != nil {
		ev := builtin_ev.NewMessageEvent(
			data.Contents,
			time.Now().Format(time.RFC3339),
			ctx.Client.GetUser(),
			data.RoomID,
		)

		ctx.Broadcast(
			builtin_ev.Generic(ev).ToRoom(*data.RoomID),
		)
	} else {
		return fmt.Errorf("TODO: global messages")
	}

	return nil
}
