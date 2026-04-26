package event

import (
	"github.com/Nykenik24/enkrypted/internal/id"
)

type Event struct {
	Kind     *EventKind `json:"kind"`
	ID       *id.ID     `json:"id"`
	ReplyTo  *id.ID     `json:"replyTo,omitempty"`
	Data     *EventData `json:"data"`
	Target   *Target    `json:"target,omitempty"`
	AuthInfo *AuthInfo  `json:"auth,omitempty"`
}
