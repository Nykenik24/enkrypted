package handlers

import (
	"fmt"
	"time"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var ConnectEventKind = buildKind(WebsocketNamespace, "connect")

type ConnectEvent struct {
	Timestamp string `json:"timestamp"`
	ID        uint64 `json:"userId"`
}

func NewConnectEvent(id uint64) *ConnectEvent {
	return &ConnectEvent{
		Timestamp: time.Now().Format(time.RFC3339),
		ID:        id,
	}
}

func (ev *ConnectEvent) Data() *event.EventData {
	return &event.EventData{
		"timestamp": ev.Timestamp,
		"userId":    ev.ID,
	}
}

func (ev *ConnectEvent) Kind() *event.EventKind {
	return ConnectEventKind
}

type ConnectHandler struct{}

func (h *ConnectHandler) Handle(ctx *ws.Context) error {
	BroadcastMessage(ctx, fmt.Sprintf("user joined %s", ctx.Client.GetUser().Username))
	return nil
}
