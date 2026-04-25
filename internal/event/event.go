package event

type Event struct {
	Kind     *EventKind `json:"kind"`
	Data     *EventData `json:"data"`
	Target   *Target    `json:"target,omitempty"`
	AuthInfo *AuthInfo  `json:"auth,omitempty"`
}
