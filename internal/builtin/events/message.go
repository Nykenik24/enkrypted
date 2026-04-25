package builtin_ev

import (
	"fmt"
	"log"
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
	RoomID    *uint64    `json:"roomId,omitempty"`
}

func NewMessageEvent(contents, timestamp string, user *user.User, room *uint64) *MessageEvent {
	lastMessageID++
	return &MessageEvent{
		Contents:  contents,
		Timestamp: timestamp,
		User:      user,
		ID:        lastMessageID,
		RoomID:    room,
	}
}

func ValidateMessage(msg *MessageEvent) error {
	if msg.ID < lastMessageID {
		return fmt.Errorf("message's ID=%d is less than last message's ID=%d", msg.ID, lastMessageID)
	}

	lastMessageID++

	log.Printf("message validated id=%d, lastId=%d", msg.ID, lastMessageID)

	return nil
}

func BroadcastMessage(ctx *ws.Context, contents string) {
	ctx.Broadcast(Generic(
		NewMessageEvent(
			contents,
			time.Now().Format(time.RFC3339),
			ctx.Client.GetUser(),
			nil,
		),
	))
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
