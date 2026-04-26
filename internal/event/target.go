package event

import "github.com/Nykenik24/enkrypted/internal/id"

type Target struct {
	RoomID    *id.ID `json:"roomId,omitempty"`
	UserID    *id.ID `json:"userId,omitempty"`
	Broadcast bool   `json:"broadcast,omitempty"`
}

func (e *Event) ToRoom(id *id.ID) *Event {
	e.Target = &Target{RoomID: id}
	return e
}

func (e *Event) ToUser(id *id.ID) *Event {
	e.Target = &Target{UserID: id}
	return e
}

func (e *Event) ToAll() *Event {
	e.Target = &Target{Broadcast: true}
	return e
}
