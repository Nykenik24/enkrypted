package handlers

import (
	"time"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var MessageEventKind = buildKind(RoomNamespace, "message")

type MessageEvent struct {
	Contents  string     `json:"contents"`
	Timestamp string     `json:"timestamp"`
	User      *user.User `json:"user"`
	ID        uint64     `json:"id"`
}

func NewMessageEvent(contents, timestamp string, user *user.User) *MessageEvent {
	lastMessageID++
	return &MessageEvent{
		Contents:  contents,
		Timestamp: timestamp,
		User:      user,
		ID:        lastMessageID,
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

func (h *MessageHandler) Handle(ctx *ws.Context) error {
	var msg struct {
		RoomID   uint64 `json:"roomId"`
		Contents string `json:"contents"`
	}

	if err := ctx.BindData(&msg); err != nil {
		return err
	}

	ev := NewMessageEvent(
		msg.Contents,
		time.Now().Format(time.RFC3339),
		ctx.Client.GetUser(),
	)

	ctx.Broadcast(
		Base(ev).ToRoom(msg.RoomID),
	)

	return nil
}
