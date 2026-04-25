package event

import "fmt"

func (e *Event) WithPassword(hash string) *Event {
	e.AuthInfo = &AuthInfo{Password: hash}
	return e
}

func (e *Event) HasAuth() bool {
	return e.AuthInfo != nil
}

func (e *Event) GetPassword() (string, error) {
	if !e.HasAuth() {
		return "", fmt.Errorf("event has no auth information")
	}

	return e.AuthInfo.Password, nil
}
