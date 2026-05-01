package user

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/Nykenik24/enkrypted/internal/id"
)

type User struct {
	Username string `json:"username"`
	ID       *id.ID `json:"id"`
}

func NewUser(username string) *User {
	return &User{
		Username: username,
		ID:       id.RandomID(),
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
