package middleware

import "github.com/Nykenik24/enkrypted/internal/event"

type Middleware interface {
	Inject(*event.Event) (*event.Event, error)
	GetName() string
}
