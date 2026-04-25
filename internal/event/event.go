package event

import (
	"math/rand"
	"strconv"
)

type Event struct {
	Kind     *EventKind `json:"kind"`
	ID       string     `json:"id"`
	ReplyTo  *string    `json:"replyTo,omitempty"`
	Data     *EventData `json:"data"`
	Target   *Target    `json:"target,omitempty"`
	AuthInfo *AuthInfo  `json:"auth,omitempty"`
}

func GenerateID() string {
	idn := rand.Uint64()
	return strconv.FormatUint(idn, 10)
}
