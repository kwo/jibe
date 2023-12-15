package jibe

import (
	"context"
	"testing"
)

func TestGetSet(t *testing.T) {
	ctx := context.Background()
	if GetID(ctx) != "" {
		t.Error("GetID(ctx) != ")
	}
	ctx = SetID(ctx, "test")
	if GetID(ctx) != "test" {
		t.Error("GetID(ctx) != test")
	}
}
