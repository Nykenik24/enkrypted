package ws

import (
	"encoding/json"
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/models"
)

type EventHandler func(ev *Event) (response *Event)

type Event struct {
	Kind string     `json:"kind"`
	ID   *models.ID `json:"id"`

	Error *string `json:"error,omitempty"`
}

func NewEvent(kind string) *Event {
	return &Event{
		Kind: kind,
		ID:   models.RandomID(),
	}
}

func UnmarshalEvent(rawJSON []byte) (*Event, error) {
	var ev Event
	err := json.Unmarshal(rawJSON, &ev)
	if err != nil {
		return nil, err
	}

	return &ev, err
}

func (ev *Event) JSON() (rawJSON []byte, err error) {
	rawJSON, err = json.Marshal(ev)
	if err != nil {
		return nil, err
	}

	return rawJSON, nil
}

func (ev *Event) String() string {
	return fmt.Sprintf(`Event:
  kind=%s
  id=%s`, ev.Kind, ev.ID)
}

func (ev *Event) HasError() bool {
	return ev.Error != nil
}
