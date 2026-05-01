package event

type AuthInfo struct {
	Password string `json:"password,omitempty"`
}

func (e *Event) WithPassword(hash string) *Event {
	e.AuthInfo = &AuthInfo{Password: hash}
	return e
}

func (e *Event) HasAuth() bool {
	return e.AuthInfo != nil
}

func (e *Event) GetPassword() string {
	return e.AuthInfo.Password
}
