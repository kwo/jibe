package jibe

import (
	"context"
	"log/slog"
)

// WithSlogHandler returns a slog.Handler that extracts the ID from the context
// and adds it to the record.
func WithSlogHandler(idName string, h slog.Handler) slog.Handler {
	return &idHandler{handler: h, idName: idName}
}

type idHandler struct {
	idName  string
	handler slog.Handler
}

func (h *idHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return h.handler.Enabled(ctx, l)
}

func (h *idHandler) Handle(ctx context.Context, r slog.Record) error {
	if id := GetID(ctx); id != "" {
		r.Add(h.idName, slog.StringValue(id))
	}
	return h.handler.Handle(ctx, r)
}

// WithAttrs implements slog.Handler.
func (h *idHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &idHandler{
		idName:  h.idName,
		handler: h.handler.WithAttrs(attrs),
	}
}

// WithGroup implements slog.Handler.
func (h *idHandler) WithGroup(name string) slog.Handler {
	return &idHandler{
		idName:  h.idName,
		handler: h.handler.WithGroup(name),
	}
}
