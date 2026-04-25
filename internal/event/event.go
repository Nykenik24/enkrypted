package event

type Target struct {
	RoomID    *uint64 `json:"roomId,omitempty"`
	UserID    *uint64 `json:"userId,omitempty"`
	Broadcast bool    `json:"broadcast,omitempty"`
}

type AuthInfo struct {
	HashedPassword string `json:"password_hash,omitempty"`
}

type Event struct {
	Kind     *EventKind `json:"kind"`
	Data     *EventData `json:"data"`
	Target   *Target    `json:"target,omitempty"`
	AuthInfo *AuthInfo  `json:"auth,omitempty"`
}
