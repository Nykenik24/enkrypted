package user

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

var lastID uint64 = 0

type User struct {
	Username string `json:"username"`
	ID       uint64 `json:"id"`
}

func NewUser(username string) *User {
	lastID++
	return &User{
		Username: username,
		ID:       lastID,
	}
}

func (u *User) Marshal() ([]byte, error) {
	rawJSON, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return rawJSON, nil
}

func Unmarshal(rawJSON []byte) (*User, error) {
	var user User
	err := json.Unmarshal(rawJSON, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GenericUsername() string {
	return fmt.Sprintf("guest_%d", rand.Uint32())
}
