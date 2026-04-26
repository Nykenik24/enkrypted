package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var LeaveRoomEventKind = buildKind(RoomNamespace, "leave")

type LeaveRoomEvent struct {
	RoomID *id.ID `json:"roomId"`
}

func NewLeaveRoomEvent(roomId *id.ID) *LeaveRoomEvent {
	return &LeaveRoomEvent{RoomID: roomId}
}

func (ev *LeaveRoomEvent) Data() *event.EventData {
	return &event.EventData{"roomId": ev.RoomID}
}

func (ev *LeaveRoomEvent) Kind() *event.EventKind {
	return LeaveRoomEventKind
}

type LeaveRoomHandler struct{}

func (h *LeaveRoomHandler) Handle(ctx *ws.Context) (*event.Event, error) {
	var ev LeaveRoomEvent

	if err := ctx.BindData(&ev); err != nil {
		return nil, err
	}

	room, err := ctx.Server.GetRoom(ev.RoomID)
	if err != nil {
		return nil, err
	}

	room.Leave(ctx.Client)

	return nil, nil
}
