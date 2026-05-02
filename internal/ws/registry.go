package ws

import "slices"

const (
	ConnectEvent    = "connect"
	DisconnectEvent = "disconnect"
	ErrorEvent      = "error"
)

var eventKinds = []string{ConnectEvent, DisconnectEvent, ErrorEvent}

func LookupKind(kind string) bool {
	return slices.Contains(eventKinds, kind)
}
