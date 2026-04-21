package event

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type EventKind int

const (
	MessageEvent EventKind = iota
	IdentifyEvent
	HandshakeStartEvent
	HandshakeReplyEvent
)

var eventKindToString = map[EventKind]string{
	MessageEvent:  "enkr:comm:msg",
	IdentifyEvent: "enkr:room:identify",
}

func stringToEventKind(str string) (EventKind, error) {
	for k, v := range eventKindToString {
		if v == str {
			return k, nil
		}
	}

	return 0, fmt.Errorf("Event '%s' doesn't exist", str)
}

var kindStringRegex = regexp.MustCompile(`^[A-Za-z0-9_-]+:[A-Za-z0-9_-]+:[A-Za-z0-9_-]+$`)

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

	v, err := stringToEventKind(s)
	if err != nil {
		return err
	}

	*k = v
	return nil
}

type Event struct {
	Kind    EventKind `json:"kind"`
	Payload any       `json:"payload"`
}

func NewEvent(kind EventKind, payload any) *Event {
	return &Event{
		Kind:    kind,
		Payload: payload,
	}
}

func (ev *Event) Marshal() ([]byte, error) {
	rawJSON, err := json.MarshalIndent(ev, "", "\t")
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
