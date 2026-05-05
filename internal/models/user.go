package models

type User struct {
	IDModel
	Connected bool `json:"connected"`

	Username string `json:"username"`
	Password string `json:"-"`
}
