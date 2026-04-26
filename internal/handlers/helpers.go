package handlers

import (
	"encoding/json"
	"time"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

func ToEventData(data any) *event.EventData {
	if data == nil {
		ed := event.EventData{}
		return &ed
	}

	if ed, ok := data.(event.EventData); ok {
		return &ed
	}
	if ed, ok := data.(*event.EventData); ok {
		return ed
	}

	raw, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil
	}

	ed := event.EventData(m)
	return &ed
}

func BroadcastMessage(ctx *ws.Context, contents string) {
	bcast := &event.Event{
		Kind: MessageEventKind,
		Data: ToEventData(NewMessageEvent(
			contents,
			time.Now().Format(time.RFC3339),
			ctx.Client.GetUser(),
		)),
		ID: id.RandomID(),
	}

	ctx.Broadcast(bcast)
}
