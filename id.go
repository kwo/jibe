package jibe

import (
	"context"
)

type contextKey string

const (
	idContextKey contextKey = "github.com/kwo/jibe/id"
)

// GetID returns the ID from the context.
func GetID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if id, ok := ctx.Value(idContextKey).(string); ok {
		return id
	}
	return ""
}

// SetID adds the ID to the context.
func SetID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, idContextKey, id)
}
