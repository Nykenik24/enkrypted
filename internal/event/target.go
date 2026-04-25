package event

func (e *Event) ToRoom(id uint64) *Event {
	e.Target = &Target{RoomID: &id}
	return e
}

func (e *Event) ToUser(id uint64) *Event {
	e.Target = &Target{UserID: &id}
	return e
}

func (e *Event) ToAll() *Event {
	e.Target = &Target{Broadcast: true}
	return e
}
