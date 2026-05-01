package enkrypt

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

type PrettyHandler struct {
	h slog.Handler
}

func (p *PrettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return p.h.Enabled(ctx, level)
}

func (p *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()
	time := r.Time.Format("15:04:05")

	var attrs []string
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, fmt.Sprintf("%s=%v", a.Key, a.Value))
		return true
	})

	fmt.Printf("[%s] %-5s %s", time, level, r.Message)
	if len(attrs) > 0 {
		fmt.Printf(" | %s", strings.Join(attrs, " "))
	}
	fmt.Println()

	return nil
}

func (p *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{h: p.h.WithAttrs(attrs)}
}

func (p *PrettyHandler) WithGroup(name string) slog.Handler {
	return &PrettyHandler{h: p.h.WithGroup(name)}
}
