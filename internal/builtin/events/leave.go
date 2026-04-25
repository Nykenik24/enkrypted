package builtin_ev

import "github.com/Nykenik24/enkrypted/internal/event"

var LeaveRoomEventKind = buildKind(RoomNamespace, "leave")

type LeaveRoomEvent struct {
	RoomID uint64 `json:"roomId"`
}

func NewLeaveRoomEvent(roomId uint64) *LeaveRoomEvent {
	return &LeaveRoomEvent{RoomID: roomId}
}

func (ev *LeaveRoomEvent) Data() *event.EventData {
	return &event.EventData{"roomId": ev.RoomID}
}

func (ev *LeaveRoomEvent) Kind() *event.EventKind {
	return LeaveRoomEventKind
}
