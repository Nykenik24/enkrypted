package event

import (
	"encoding/json"
	"fmt"
	"log"
)

type eventJSON struct {
	Kind string         `json:"kind"`
	Data map[string]any `json:"data"`
	Auth *AuthInfo      `json:"auth,omitempty"`
}

func EventFromJSON(rawJSON []byte) (*Event, error) {
	var evJSON eventJSON
	log.Printf("unmarshaling raw JSON into eventJSON: %s", rawJSON)
	if err := json.Unmarshal(rawJSON, &evJSON); err != nil {
		return nil, err
	}

	log.Printf("getting kind from string '%s'", evJSON.Kind)
	kind, err := KindFromString(evJSON.Kind)
	if err != nil {
		return nil, fmt.Errorf("KindFromString: %s", err)
	}

	ev := &Event{
		Data:     (*EventData)(&evJSON.Data),
		Kind:     kind,
		AuthInfo: evJSON.Auth,
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
