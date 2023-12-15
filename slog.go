package jibe

import (
	"context"
	"log/slog"
)

// SlogHandler returns a slog.Handler that extracts the ID from the context
// and adds it to the record.
func SlogHandler(h slog.Handler, idName string) slog.Handler {
	return idHandler{Handler: h, idName: idName}
}

type idHandler struct {
	slog.Handler
	idName string
}

func (h idHandler) Handle(ctx context.Context, r slog.Record) error {
	if id := GetID(ctx); id != "" {
		r.Add(h.idName, slog.StringValue(id))
	}
	return h.Handler.Handle(ctx, r)
}
