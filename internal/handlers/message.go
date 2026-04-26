package handlers

import (
	"time"

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

func (ev *MessageEvent) Data() *event.EventData {
	return &event.EventData{
		"contents":  ev.Contents,
		"timestamp": ev.Timestamp,
		"user":      ev.User,
		"id":        ev.ID,
	}
}

func (ev *MessageEvent) Kind() *event.EventKind {
	return MessageEventKind
}

type MessageHandler struct{}

func (h *MessageHandler) Handle(ctx *ws.Context) (*event.Event, error) {
	var data MessageEvent

	if err := ctx.BindData(&data); err != nil {
		return nil, err
	}

	ev := NewMessageEvent(
		data.Contents,
		time.Now().Format(time.RFC3339),
		ctx.Client.GetUser(),
	)

	ctx.Broadcast(
		Base(ev).ToRoom(data.RoomID),
	)

	return nil, nil
}
