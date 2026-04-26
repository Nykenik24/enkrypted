package handlers

import (
	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var (
	GetRoomsEventKind = buildKind(RoomNamespace, "all")
	GetRoomsReplyKind = buildKind(RoomNamespace, "all_reply")
)

type GetRoomsEvent struct{}

func NewGetRoomsEvent() *GetRoomsEvent {
	return &GetRoomsEvent{}
}

func (ev *GetRoomsEvent) Data() *event.EventData {
	return &event.EventData{}
}

func (ev *GetRoomsEvent) Kind() *event.EventKind {
	return GetRoomsEventKind
}

type GetRoomsReply struct {
	Count uint32    `json:"count"`
	Rooms []ws.Room `json:"rooms"`
}

func NewGetRoomsReply(rooms []ws.Room) *GetRoomsReply {
	return &GetRoomsReply{
		Count: uint32(len(rooms)),
		Rooms: rooms,
	}
}

func (ev *GetRoomsReply) Data() *event.EventData {
	return &event.EventData{
		"count": ev.Count,
		"rooms": ev.Rooms,
	}
}

func (ev *GetRoomsReply) Kind() *event.EventKind {
	return GetRoomsReplyKind
}

type GetRoomsHandler struct{}

func (h *GetRoomsHandler) Handle(ctx *ws.Context) (*event.Event, error) {
	var rooms []ws.Room

	for _, room := range ctx.Server.GetAllRooms() {
		rooms = append(rooms, room)
	}

	return Base(NewGetRoomsReply(rooms)), nil
}
