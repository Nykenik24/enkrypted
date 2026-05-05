package models

type User struct {
	IDModel
	Connected bool `json:"connected"`

	Username   string `json:"username"`
	HashedPass string `json:"hashed_password"`
}
