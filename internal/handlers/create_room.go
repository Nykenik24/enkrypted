package handlers

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/event"
	"github.com/Nykenik24/enkrypted/internal/ws"
)

var CreateRoomEventKind = buildKind(RoomNamespace, "create")

var CreateRoomHandler = BuildHandler(CreateRoomEventKind, func(ctx *ws.Context) (*event.Event, error) {
	ev := ctx.Event

	if !ev.HasAuth() {
		return nil, fmt.Errorf("%s must have auth information", CreateRoomEventKind.String())
	}

	passwd := ev.GetPassword()

	auth := ctx.Server.GetAuth()

	if !auth.Hasher.VerifyPassword(passwd, auth.GetAdminHash()) {
		return nil, fmt.Errorf("wrong password")
	}

	_, err := ctx.Server.AddRoom()
	if err != nil {
		return nil, err
	}

	BroadcastMessage(ctx, "created room")

	return nil, nil
})
