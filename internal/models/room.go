package models

import "github.com/Nykenik24/enkrypted/internal/id"

type Room struct {
	id.IDModel
	Password string `json:"-"`
}
