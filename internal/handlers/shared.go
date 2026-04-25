package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
)

var lastMessageID uint64 = 0

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

type BuiltinEvent interface {
	Data() *event.EventData
	Kind() *event.EventKind
}

func Base(ev BuiltinEvent) *event.Event {
	return &event.Event{
		Kind: ev.Kind(),
		Data: ev.Data(),
		ID:   event.GenerateID(),
	}
}
