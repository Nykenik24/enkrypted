package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var MessageEventKind = buildKind(RoomNamespace, "message")

type MessageEvent struct {
	Contents  string     `json:"contents"`
	Timestamp string     `json:"timestamp"`
	User      *user.User `json:"user"`
	ID        *id.ID     `json:"id"`
	RoomID    *id.ID     `json:"roomId"`
}

func NewMessageEvent(contents, timestamp string, user *user.User) *MessageEvent {
	return &MessageEvent{
		Contents:  contents,
		Timestamp: timestamp,
		User:      user,
		ID:        id.RandomID(),
	}
}

var MessageHandler = BuildHandler(MessageEventKind, func(ctx *ws.Context) (*event.Event, error) {
	var data MessageEvent

	if err := ctx.BindData(&data); err != nil {
		return nil, err
	}

	ev := &event.Event{
		Kind: MessageEventKind,
		Data: ToEventData(data),
	}
	ctx.Broadcast(
		ev.ToRoom(data.RoomID),
	)

	return nil, nil
})
