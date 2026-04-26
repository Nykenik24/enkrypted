package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

type baseHandler struct {
	handleFunc HandleFunc
	kind       *event.EventKind
}

func (h *baseHandler) Handle(ctx *ws.Context) (*event.Event, error) {
	reply, err := h.handleFunc(ctx)

	if reply != nil {
		reply.ID = id.RandomID()
	}

	return reply, err
}

func (h *baseHandler) Kind() *event.EventKind {
	return h.kind
}

func BuildHandler(kind *event.EventKind, handleFunc HandleFunc) *baseHandler {
	h := &baseHandler{
		handleFunc: handleFunc,
		kind:       kind,
	}

	return h
}
