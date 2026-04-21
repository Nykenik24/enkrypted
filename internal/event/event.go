package event

import (
	"encoding/json"
	"fmt"
	"log"
)

type EventData map[string]any

func (ed *EventData) JSON() ([]byte, error) {
	rawJSON, err := json.Marshal(ed)
	if err != nil {
		return nil, err
	}

	return rawJSON, nil
}

type Event struct {
	Kind *EventKind `json:"kind"`
	Data *EventData `json:"data"`
}

type eventJSON struct {
	Kind string         `json:"kind"`
	Data map[string]any `json:"data"`
}

func EventFromJSON(rawJSON []byte) (*Event, error) {
	var evJSON eventJSON
	log.Println("Unmarshaling raw JSON into eventJSON")
	if err := json.Unmarshal(rawJSON, &evJSON); err != nil {
		return nil, err
	}

	log.Printf("Getting kind from string '%s'", evJSON.Kind)
	kind, err := KindFromString(evJSON.Kind)
	if err != nil {
		return nil, fmt.Errorf("KindFromString: %s", err)
	}

	log.Println("Building event")
	ev := &Event{
		Data: (*EventData)(&evJSON.Data),
		Kind: kind,
	}

	return ev, nil
}

func (ev *Event) JSON() ([]byte, error) {
	rawJSON, err := json.MarshalIndent(ev, "", "  ")
	if err != nil {
		return nil, err
	}

	return rawJSON, nil
}
