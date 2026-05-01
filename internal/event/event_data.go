package event

import "encoding/json"

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
