package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var CreateRoomEventKind = buildKind(RoomNamespace, "create")

type CreateRoomEvent struct{}

func NewCreateRoomEvent() *CreateRoomEvent {
	return &CreateRoomEvent{}
}

func (ev *CreateRoomEvent) Data() *event.EventData {
	return &event.EventData{}
}

func (ev *CreateRoomEvent) Kind() *event.EventKind {
	return CreateRoomEventKind
}

type CreateRoomHandler struct{}

func (h *CreateRoomHandler) Handle(ctx *ws.Context) (*event.Event, error) {
	ev := ctx.Event

	if !ev.HasAuth() {
		return nil, fmt.Errorf("%s must have auth information", CreateRoomEventKind.String())
	}

	passwd, err := ev.GetPassword()
	if err != nil {
		return nil, err
	}

	auth := ctx.Server.GetAuth()

	if !auth.Hasher.VerifyPassword(passwd, auth.GetAdminHash()) {
		return nil, fmt.Errorf("wrong password")
	}

	_, err = ctx.Server.AddRoom()
	if err != nil {
		return nil, err
	}

	BroadcastMessage(ctx, "created room")

	return nil, nil
}
