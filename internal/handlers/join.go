package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var JoinRoomEventKind = buildKind(RoomNamespace, "join")

type JoinRoomEvent struct {
	RoomID *id.ID `json:"roomId"`
}

func NewJoinRoomEvent(roomId *id.ID) *JoinRoomEvent {
	return &JoinRoomEvent{RoomID: roomId}
}

var JoinRoomHandler = BuildHandler(JoinRoomEventKind, func(ctx *ws.Context) (*event.Event, error) {
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
})
