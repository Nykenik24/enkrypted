package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var RemoveRoomEventKind = buildKind(RoomNamespace, "remove")

type RemoveRoomEvent struct {
	RoomID *id.ID `json:"roomId"`
}

func NewRemoveRoomEvent(id *id.ID) *RemoveRoomEvent {
	return &RemoveRoomEvent{RoomID: id}
}

func (ev *RemoveRoomEvent) Data() *event.EventData {
	return &event.EventData{"roomId": ev.RoomID}
}

func (ev *RemoveRoomEvent) Kind() *event.EventKind {
	return RemoveRoomEventKind
}

type RemoveRoomHandler struct{}

func (h *RemoveRoomHandler) Handle(ctx *ws.Context) (*event.Event, error) {
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

	var data RemoveRoomEvent
	if err := ctx.BindData(&data); err != nil {
		return nil, err
	}

	err = ctx.Server.RemoveRoom(data.RoomID)
	if err != nil {
		return nil, err
	}

	BroadcastMessage(ctx, fmt.Sprintf("removed room %d", data.RoomID))

	return nil, nil
}
