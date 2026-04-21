package builtin_ev

import (
	"github.com/Nykenik24/enkrypted/internal/event"
)

type IdentifyEvent struct {
	Username string `json:"username"`
}

func NewIdentifyEvent(username string) *IdentifyEvent {
	return &IdentifyEvent{
		Username: username,
	}
}

func (ev *IdentifyEvent) Data() *event.EventData {
	return &event.EventData{
		"username": ev.Username,
	}
}

func (ev *IdentifyEvent) Kind() *event.EventKind {
	return event.NewEventKind(definee, "room", "identify")
}
