package event

import "github.com/Nykenik24/enkrypted/internal/id"

func (e *Event) RepliesTo(id *id.ID) *Event {
	e.ReplyTo = id
	return e
}

func (e *Event) ClearReply() *Event {
	e.ReplyTo = nil
	return e
}
