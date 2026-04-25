package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var IdentifyEventKind = buildKind(RoomNamespace, "identify")

type IdentifyEvent struct {
	Username string `json:"username"`
}

func NewIdentifyEvent(username string) *IdentifyEvent {
	return &IdentifyEvent{
		Username: username,
	}
}

func (ev *IdentifyEvent) Data() *event.EventData {
	return &event.EventData{
		"username": ev.Username,
	}
}

func (ev *IdentifyEvent) Kind() *event.EventKind {
	return IdentifyEventKind
}

type IdentifyHandler struct{}

func (h *IdentifyHandler) Handle(ctx *ws.Context) error {
	var ev IdentifyEvent

	if err := ctx.BindData(&ev); err != nil {
		return err
	}

	if ev.Username == "" {
		ev.Username = user.GenericUsername()
	}

	ctx.Client.SetUsername(ev.Username)

	BroadcastMessage(ctx, fmt.Sprintf("user joined %s", ctx.Client.GetUser().Username))
	return nil
}
