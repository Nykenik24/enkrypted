package middleware

import "github.com/Nykenik24/enkrypted/internal/event"

type MiddlewareManager struct {
	midware map[string]Middleware
}

func NewMidwareManager() *MiddlewareManager {
	return &MiddlewareManager{
		midware: make(map[string]Middleware),
	}
}

func (mm *MiddlewareManager) Register(midware Middleware) {
	mm.midware[midware.GetName()] = midware
}

func (mm *MiddlewareManager) Unregister(midware Middleware) {
	delete(mm.midware, midware.GetName())
}

func (mm *MiddlewareManager) MultiInject(ev *event.Event) (*event.Event, error) {
	var newEv *event.Event = ev
	var err error

	for _, midware := range mm.midware {
		newEv, err = midware.Inject(newEv)
		if err != nil {
			return nil, err
		}
	}

	return newEv, nil
}
