package ws

import (
	"github.com/Nykenik24/enkrypted/internal/event"
)

type Context struct {
	Client *Client
	Hub    *Hub
	Event  *event.Event
	Server Server
}

func (ctx *Context) Bind(v any) error {
	return ctx.Event.Decode(v)
}

func (ctx *Context) Broadcast(ev *event.Event) {
	ctx.Hub.Broadcast(ev)
}

func (ctx *Context) Emit(ev *event.Event) {
	ctx.Hub.Emit(ctx.Client, ev)
}
