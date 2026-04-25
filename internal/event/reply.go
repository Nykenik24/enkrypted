package event

func (e *Event) RepliesTo(id string) *Event {
	e.ReplyTo = &id
	return e
}

func (e *Event) ClearReply() *Event {
	e.ReplyTo = nil
	return e
}
