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
	Status   int    `json:"status"`
	Username string `json:"username"`
	Reason   string `json:"reason,omitempty"`
}

func NewIdentifyAck(status int, username string, reason string) *IdentifyAcknowledge {
	return &IdentifyAcknowledge{
		Status:   status,
		Username: username,
		Reason:   reason,
	}
}

const username_max_len = 32

var usernameRegex = regexp.MustCompile(`^[\w\d]+$`)

func validUsername(username string) error {
	if len(username) > username_max_len {
		return fmt.Errorf("username longer than %d characters", username_max_len)
	}

	if !usernameRegex.Match([]byte(username)) {
		return fmt.Errorf("invalid format")
	}

	return nil
}

var IdentifyHandler = BuildHandler(IdentifyEventKind, func(ctx *ws.Context) (*event.Event, error) {
	var data IdentifyEvent
	var err error

	if err := ctx.BindData(&data); err != nil {
		return nil, err
	}

	username := data.Username

	if username == "" {
		username = user.GenericUsername()
	}

	var status int = 0
	var reason string = "ok"

	if err = validUsername(username); err != nil {
		reason = err.Error()
		status = 1
	} else {
		for _, client := range ctx.Hub.GetAllClients() {
			if username == client.GetUser().Username {
				err = fmt.Errorf("username taken")
				status = 1
				break
			}
		}
	}

	if status == 0 {
		ctx.Client.SetUsername(username)
		BroadcastMessage(ctx, fmt.Sprintf("user joined %s", ctx.Client.GetUser().Username))
	}

	replyData := NewIdentifyAck(status, username, reason)

	reply := &event.Event{
		Kind: IdentifyAckKind,
		Data: ToEventData(replyData),
	}
	reply.RepliesTo(ctx.Event.ID)

	return reply, err
})

// func (h *IdentifyHandler) Handle(ctx *ws.Context) (*event.Event, error) {

// }
