package event

func (e *Event) WithPassword(hash string) *Event {
	e.AuthInfo = &AuthInfo{HashedPassword: hash}
	return e
}
