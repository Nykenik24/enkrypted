package builtin_ev

import "github.com/Nykenik24/enkrypted/internal/event"

const definee = "enkr"

const (
	MessageEventKind  = definee + ":comm:message"
	IdentifyEventKind = definee + ":room:identify"
)

type BuiltinEvent interface {
	Data() *event.EventData
	Kind() *event.EventKind
}

func ToGeneric(ev BuiltinEvent) *event.Event {
	generic := &event.Event{
		Kind: ev.Kind(),
		Data: ev.Data(),
	}

	return generic
}
