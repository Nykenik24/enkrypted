package event

import (
	"encoding/json"
	"fmt"
)

type EventKind int

const (
	PingEvent EventKind = iota
	MessageEvent
)

var eventKindToString = map[EventKind]string{
	PingEvent:    "enkr:ping",
	MessageEvent: "enkr:msg",
}

var stringToEventKind = map[string]EventKind{
	"enkr:ping": PingEvent,
	"enkr:msg":  MessageEvent,
}

func (k EventKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(eventKindToString[k])
}

func (k *EventKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	v, ok := stringToEventKind[s]
	if !ok {
		return fmt.Errorf("unknown event kind: %s", s)
	}

	*k = v
	return nil
}

type Event struct {
	Kind    EventKind       `json:"kind"`
	Payload json.RawMessage `json:"payload"`
}
