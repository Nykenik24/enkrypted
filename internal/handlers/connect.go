package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var ConnectEventKind = buildKind(WebsocketNamespace, "connect")

type ConnectEvent struct {
	Timestamp string `json:"timestamp"`
	ID        *id.ID `json:"userId"`
}

var ConnectHandler = BuildHandler(ConnectEventKind, func(ctx *ws.Context) (*event.Event, error) {
	// could add some logic here later (such as handshaking?).
	return nil, nil
})
