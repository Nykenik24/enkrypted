package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

const definee = "enkr"

type Namespace int

const (
	RoomNamespace Namespace = iota
	WebsocketNamespace
)

var namespaces = map[Namespace]string{
	RoomNamespace:      "room",
	WebsocketNamespace: "ws",
}

func buildKind(ns Namespace, name string) *event.EventKind {
	kind, _ := event.KindFromString(fmt.Sprintf("%s:%s:%s", definee, namespaces[ns], name))
	return kind
}

type HandleFunc func(*ws.Context) (*event.Event, error)
