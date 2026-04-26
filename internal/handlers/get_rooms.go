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

var GetRoomsHandler = BuildHandler(GetRoomsEventKind, func(ctx *ws.Context) (*event.Event, error) {
	var rooms []ws.Room

	for _, room := range ctx.Server.GetAllRooms() {
		rooms = append(rooms, room)
	}

	replyData := NewGetRoomsReply(rooms)

	reply := &event.Event{
		Kind: GetRoomsReplyKind,
		Data: ToEventData(replyData),
	}

	return reply, nil
})
