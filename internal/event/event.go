package event

type Target struct {
	RoomID    *uint64 `json:"roomId,omitempty"`
	UserID    *uint64 `json:"userId,omitempty"`
	Broadcast bool    `json:"broadcast,omitempty"`
}

type AuthInfo struct {
	Password string `json:"password,omitempty"`
}

type Event struct {
	Kind     *EventKind `json:"kind"`
	Data     *EventData `json:"data"`
	Target   *Target    `json:"target,omitempty"`
	AuthInfo *AuthInfo  `json:"auth,omitempty"`
}
