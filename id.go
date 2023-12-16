package jibe

import (
	"context"
)

type contextKey string

type IDProducer func() string

const (
	idContextKey contextKey = "github.com/kwo/jibe/id"
)

func GetID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if id, ok := ctx.Value(idContextKey).(string); ok {
		return id
	}
	return ""
}

func SetID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, idContextKey, id)
}
