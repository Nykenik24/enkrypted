package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var JoinRoomEventKind = buildKind(RoomNamespace, "join")
var JoinRoomAckEventKind = buildKind(RoomNamespace, "join_ack")

type JoinRoomEvent struct {
	RoomID *id.ID `json:"roomId"`
}

func NewJoinRoomEvent(roomId *id.ID) *JoinRoomEvent {
	return &JoinRoomEvent{RoomID: roomId}
}

type JoinRoomAckEvent struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
	RoomID *id.ID `json:"roomId"`
}

func NewJoinRoomAckEvent(status int, roomId *id.ID, reason string) *JoinRoomAckEvent {
	return &JoinRoomAckEvent{Status: status, RoomID: roomId, Reason: reason}
}

var JoinRoomHandler = BuildHandler(JoinRoomEventKind, func(ctx *ws.Context) (*event.Event, error) {
	var ev JoinRoomEvent
	var status int = 0
	var reason string = "ok"

	if err := ctx.BindData(&ev); err != nil {
		status = 1
		reason = err.Error()
	}

	room, err := ctx.Server.GetRoom(ev.RoomID)
	if err != nil {
		status = 1
		reason = err.Error()
	}

	if room == nil {
		status = 1
		reason = "room not found"
	} else {
		if err := room.Join(ctx.Client); err != nil {
			status = 1
			reason = err.Error()
		}
	}

	replyData := NewJoinRoomAckEvent(status, ev.RoomID, reason)

	reply := &event.Event{
		Kind: JoinRoomAckEventKind,
		Data: ToEventData(replyData),
	}
	reply.RepliesTo(ctx.Event.ID)

	return reply, err
})
