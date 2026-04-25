package builtin_ev

import (
	"github.com/Nykenik24/enkrypted/internal/event"
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
