package builtin_ev

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/user"
)

var lastID uint64 = 0

type MessageEvent struct {
	Contents  string     `json:"contents"`
	Timestamp string     `json:"timestamp"`
	User      *user.User `json:"user"`
	ID        uint64     `json:"id"`
}

func NewMessageEvent(contents, timestamp string, user *user.User) *MessageEvent {
	lastID++
	return &MessageEvent{
		Contents:  contents,
		Timestamp: timestamp,
		User:      user,
		ID:        lastID,
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
	return event.NewEventKind(definee, "comm", "message")
}
