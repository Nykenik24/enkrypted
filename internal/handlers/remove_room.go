package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var RemoveRoomEventKind = buildKind(RoomNamespace, "remove")

type RemoveRoomEvent struct {
	RoomID uint64 `json:"roomId"`
}

func NewRemoveRoomEvent(id uint64) *RemoveRoomEvent {
	return &RemoveRoomEvent{RoomID: id}
}

func (ev *RemoveRoomEvent) Data() *event.EventData {
	return &event.EventData{"roomId": ev.RoomID}
}

func (ev *RemoveRoomEvent) Kind() *event.EventKind {
	return RemoveRoomEventKind
}

type RemoveRoomHandler struct{}

func (h *RemoveRoomHandler) Handle(ctx *ws.Context) error {
	ev := ctx.Event

	if !ev.HasAuth() {
		return fmt.Errorf("%s must have auth information", CreateRoomEventKind.String())
	}

	passwd, err := ev.GetPassword()
	if err != nil {
		return err
	}

	auth := ctx.Server.GetAuth()
	if !auth.Hasher.VerifyPassword(passwd, auth.GetAdminHash()) {
		return fmt.Errorf("wrong password")
	}

	var data RemoveRoomEvent
	if err := ctx.BindData(&data); err != nil {
		return err
	}

	err = ctx.Server.RemoveRoom(data.RoomID)
	if err != nil {
		return err
	}

	BroadcastMessage(ctx, fmt.Sprintf("removed room %d", data.RoomID))

	return nil
}
