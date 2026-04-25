package event

import (
	"encoding/json"
	"fmt"
	"log"
)

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
