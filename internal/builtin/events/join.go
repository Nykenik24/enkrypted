package builtin_ev

import "github.com/Nykenik24/enkrypted/internal/event"

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
