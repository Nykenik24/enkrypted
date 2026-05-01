package models

type Room struct {
	StringIDModel
	Password string `json:"-"`
}
