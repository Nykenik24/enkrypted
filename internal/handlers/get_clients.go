package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/id"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var (
	GetClientsEventKind = buildKind(WebsocketNamespace, "clients")
	GetClientsReplyKind = buildKind(WebsocketNamespace, "clients_reply")
)

type GetClientsEvent struct{}

func NewGetClientsEvent() *GetClientsEvent {
	return &GetClientsEvent{}
}

type ClientData struct {
	ID       *id.ID `json:"id"`
	Username string `json:"username"`
}

type GetClientsReply struct {
	Count   uint32        `json:"count"`
	Clients []*ClientData `json:"clients"`
}

func NewGetClientsReply(clients []*ClientData) *GetClientsReply {
	return &GetClientsReply{
		Count:   uint32(len(clients)),
		Clients: clients,
	}
}

var GetClientsHandler = BuildHandler(GetClientsEventKind, func(ctx *ws.Context) (*event.Event, error) {
	var clients []*ClientData

	for _, client := range ctx.Hub.GetAllClients() {
		clients = append(clients, &ClientData{ID: client.GetID(), Username: client.GetUser().Username})
	}

	replyData := NewGetClientsReply(clients)

	reply := &event.Event{
		Kind: GetClientsReplyKind,
		Data: ToEventData(replyData),
	}

	return reply, nil
})
