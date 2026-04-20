package event

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type EventKind int

const (
	PingEvent EventKind = iota
	MessageEvent
	IdentifyEvent
)

var eventKindToString = map[EventKind]string{
	MessageEvent:  "enkr:msg",
	IdentifyEvent: "enkr:identify",
}

var stringToEventKind = map[string]EventKind{
	"enkr:msg":      MessageEvent,
	"enkr:identify": IdentifyEvent,
}

var kindStringRegex = regexp.MustCompile(`^[A-Za-z0-9_-]+(?::[A-Za-z0-9_-]+)*$`)

func validateKindString(s string) error {
	if !kindStringRegex.Match([]byte(s)) {
		return fmt.Errorf("event kind string doesn't follow the correct format")
	}
	return nil
}

func (k EventKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(eventKindToString[k])
}

func (k *EventKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if err := validateKindString(s); err != nil {
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

func (ev *Event) Marshal() ([]byte, error) {
	rawJSON, err := json.Marshal(ev)
	if err != nil {
		return nil, err
	}

	return rawJSON, nil
}

func Unmarshal(rawJSON []byte) (*Event, error) {
	var ev Event
	if err := json.Unmarshal(rawJSON, &ev); err != nil {
		return nil, err
	}

	return &ev, nil
}
