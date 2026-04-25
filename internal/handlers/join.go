package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var JoinRoomEventKind = buildKind(RoomNamespace, "join")

type JoinRoomEvent struct {
	RoomID uint64 `json:"roomId"`
}

func NewJoinRoomEvent(roomId uint64) *JoinRoomEvent {
	return &JoinRoomEvent{RoomID: roomId}
}

func (ev *JoinRoomEvent) Data() *event.EventData {
	return &event.EventData{"roomId": ev.RoomID}
}

func (ev *JoinRoomEvent) Kind() *event.EventKind {
	return JoinRoomEventKind
}

type JoinRoomHandler struct{}

func (h *JoinRoomHandler) Handle(ctx *ws.Context) (*event.Event, error) {
	var ev JoinRoomEvent

	if err := ctx.BindData(&ev); err != nil {
		return nil, err
	}

	room, err := ctx.Server.GetRoom(ev.RoomID)
	if err != nil {
		return nil, err
	}

	if err := room.Join(ctx.Client); err != nil {
		return nil, err
	}

	return nil, nil
}
