package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var (
	IdentifyEventKind = buildKind(RoomNamespace, "identify")
	IdentifyAckKind   = buildKind(RoomNamespace, "identify_ack")
)

type IdentifyEvent struct {
	Username string `json:"username"`
}

func NewIdentifyEvent(username string) *IdentifyEvent {
	return &IdentifyEvent{
		Username: username,
	}
}

type IdentifyAcknowledge struct {
	Status int `json:"status"`
}

func NewIdentifyAck(status int) *IdentifyAcknowledge {
	return &IdentifyAcknowledge{
		Status: status,
	}
}

var IdentifyHandler = BuildHandler(IdentifyEventKind, func(ctx *ws.Context) (*event.Event, error) {
	var data IdentifyEvent

	if err := ctx.BindData(&data); err != nil {
		return nil, err
	}

	if data.Username == "" {
		data.Username = user.GenericUsername()
	}

	ctx.Client.SetUsername(data.Username)

	BroadcastMessage(ctx, fmt.Sprintf("user joined %s", ctx.Client.GetUser().Username))

	reply := &event.Event{
		Kind: IdentifyAckKind,
		Data: ToEventData(data),
	}
	reply.RepliesTo(ctx.Event.ID)

	return reply, nil
})

// func (h *IdentifyHandler) Handle(ctx *ws.Context) (*event.Event, error) {

// }
