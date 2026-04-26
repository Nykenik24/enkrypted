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

func (ev *GetClientsEvent) Data() *event.EventData {
	return &event.EventData{}
}

func (ev *GetClientsEvent) Kind() *event.EventKind {
	return GetClientsEventKind
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

func (ev *GetClientsReply) Data() *event.EventData {
	return &event.EventData{
		"count":   ev.Count,
		"clients": ev.Clients,
	}
}

func (ev *GetClientsReply) Kind() *event.EventKind {
	return GetClientsReplyKind
}

type GetClientsHandler struct{}

func (h *GetClientsHandler) Handle(ctx *ws.Context) (*event.Event, error) {
	var clients []*ClientData

	for _, client := range ctx.Hub.GetAllClients() {
		clients = append(clients, &ClientData{ID: client.GetID(), Username: client.GetUser().Username})
	}

	return Base(NewGetClientsReply(clients)), nil
}
