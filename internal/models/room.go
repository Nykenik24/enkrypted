package models

type Room struct {
	IDModel
	Password string `json:"-"`
}
