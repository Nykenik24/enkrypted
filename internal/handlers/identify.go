package handlers

import (
	"fmt"
	"regexp"

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
	Status   int     `json:"status"`
	Username string  `json:"username"`
	Reason   *string `json:"reason,omitempty"`
}

func NewIdentifyAck(status int, username string) *IdentifyAcknowledge {
	return &IdentifyAcknowledge{
		Status:   status,
		Username: username,
		Reason:   nil,
	}
}

func (ack *IdentifyAcknowledge) SetReason(reason string) *IdentifyAcknowledge {
	ack.Reason = &reason
	return ack
}

const username_max_len = 32

var usernameRegex = regexp.MustCompile(`^[\w\d]+$`)

func validUsername(username string) bool {
	if len(username) > username_max_len {
		return false
	}

	return usernameRegex.Match([]byte(username))
}

var IdentifyHandler = BuildHandler(IdentifyEventKind, func(ctx *ws.Context) (*event.Event, error) {
	var data IdentifyEvent

	if err := ctx.BindData(&data); err != nil {
		return nil, err
	}

	username := data.Username

	if username == "" {
		username = user.GenericUsername()
	}

	var status int = 0
	var reason string = ""
	var err error = nil

	if validUsername(username) {
		for _, client := range ctx.Hub.GetAllClients() {
			if username == client.GetUser().Username {
				err = fmt.Errorf("username taken: %s", username)
				status = 1
				reason = "username taken"
				break
			}
		}
	} else {
		err = fmt.Errorf("invalid username: %s", username)
		status = 1
		reason = "invalid username"
	}

	if status == 0 {
		ctx.Client.SetUsername(username)
		BroadcastMessage(ctx, fmt.Sprintf("user joined %s", ctx.Client.GetUser().Username))
	}

	replyData := NewIdentifyAck(status, username)
	if status == 1 {
		replyData.SetReason(reason)
	}

	reply := &event.Event{
		Kind: IdentifyAckKind,
		Data: ToEventData(replyData),
	}
	reply.RepliesTo(ctx.Event.ID)

	return reply, err
})

// func (h *IdentifyHandler) Handle(ctx *ws.Context) (*event.Event, error) {

// }
