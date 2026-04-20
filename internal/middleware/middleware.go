package middleware

import "github.com/Nykenik24/enkrypted/internal/event"

type MiddlewareInfo struct {
	Name string
}

type Middleware interface {
	Inject(*event.Event) (*event.Event, error)
	Info() MiddlewareInfo
}
