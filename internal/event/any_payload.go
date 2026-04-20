package event

import "encoding/json"

type AnyPayloadEvent struct {
	Kind    EventKind
	Payload any

	rawJSON []byte
}

func NewAnyPayload(kind EventKind, payload any) (*AnyPayloadEvent, error) {
	anyev := AnyPayloadEvent{
		Kind:    kind,
		Payload: payload,
	}

	rawJSON, err := json.Marshal(anyev)
	if err != nil {
		return nil, err
	}

	anyev.rawJSON = rawJSON
	return &anyev, nil
}

func (anyev *AnyPayloadEvent) JSON() []byte {
	return anyev.rawJSON
}

func (anyev *AnyPayloadEvent) RegularEvent() (*Event, error) {
	ev, err := Unmarshal(anyev.rawJSON)
	if err != nil {
		return nil, err
	}

	return ev, nil
}
