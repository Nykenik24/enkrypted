package middleware

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
)

type MiddlewareInfo struct {
	Name string

	AllKinds bool
	Kinds    []event.EventKind
}

type Middleware interface {
	Inject(*event.Event) (*event.Event, error)
	Info() MiddlewareInfo
}

func MultiInject(midwares map[string]Middleware, ev *event.Event) error {
	var err error = nil
	for _, midware := range midwares {
		if midware.Info().AllKinds {
			ev, err = midware.Inject(ev)
			if err != nil {
				return fmt.Errorf("when running middleware: %v", err)
			}
		} else {
			for _, v := range midware.Info().Kinds {
				if ev.Kind == v {
					ev, err = midware.Inject(ev)
					if err != nil {
						return fmt.Errorf("when running middleware: %v", err)
					}
				}
			}
		}
	}

	return nil
}
