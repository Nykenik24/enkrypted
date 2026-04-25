package event

import (
	"encoding/json"
	"fmt"
	"log"
)

type EventData map[string]any

func (ed *EventData) JSON() ([]byte, error) {
	rawJSON, err := json.MarshalIndent(ed, "", "  ")
	if err != nil {
		return nil, err
	}

	return rawJSON, nil
}

func (ed *EventData) Get(key string) any {
	return (map[string]any)(*ed)[key]
}

type Target struct {
	RoomID    *uint64 `json:"roomId,omitempty"`
	UserID    *uint64 `json:"userId,omitempty"`
	Broadcast bool    `json:"broadcast,omitempty"`
}

type Event struct {
	Kind   *EventKind `json:"kind"`
	Data   *EventData `json:"data"`
	Target *Target    `json:"target,omitempty"`
}

type eventJSON struct {
	Kind string         `json:"kind"`
	Data map[string]any `json:"data"`
}

func EventFromJSON(rawJSON []byte) (*Event, error) {
	var evJSON eventJSON
	log.Println("unmarshaling raw JSON into eventJSON")
	if err := json.Unmarshal(rawJSON, &evJSON); err != nil {
		return nil, err
	}

	log.Printf("getting kind from string '%s'", evJSON.Kind)
	kind, err := KindFromString(evJSON.Kind)
	if err != nil {
		return nil, fmt.Errorf("KindFromString: %s", err)
	}

	log.Println("building event")
	ev := &Event{
		Data: (*EventData)(&evJSON.Data),
		Kind: kind,
	}

	return ev, nil
}

func (e *Event) Decode(v any) error {
	raw, err := json.Marshal(e.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, v)
}

func (ev *Event) JSON() ([]byte, error) {
	rawJSON, err := json.MarshalIndent(ev, "", "  ")
	if err != nil {
		return nil, err
	}

	return rawJSON, nil
}

func (e *Event) ToRoom(id uint64) *Event {
	e.Target = &Target{RoomID: &id}
	return e
}

func (e *Event) ToUser(id uint64) *Event {
	e.Target = &Target{UserID: &id}
	return e
}

func (e *Event) ToAll() *Event {
	e.Target = &Target{Broadcast: true}
	return e
}
