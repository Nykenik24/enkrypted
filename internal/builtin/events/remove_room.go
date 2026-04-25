package builtin_ev

import (
	"github.com/Nykenik24/enkrypted/internal/event"
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
