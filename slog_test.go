package jibe

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"
)

func TestSlogJson(t *testing.T) {
	buf := &bytes.Buffer{}
	var handler slog.Handler
	handler = slog.NewJSONHandler(buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	handler = WithRecordID("traceId", handler)
	slog.SetDefault(slog.New(handler))

	ctx1 := context.Background()
	ctx1 = SetID(ctx1, "1234")

	slog.InfoContext(ctx1, "test")
	if buf.String() == "" {
		t.Error("buf.String() == ")
	}
	if !strings.Contains(buf.String(), `"traceId":"1234"`) {
		t.Logf("buf.String() == %s", buf.String())
		t.Error(`!strings.Contains(buf.String(), "traceId":"1234")`)
	}
}

func TestSlogText(t *testing.T) {
	buf := &bytes.Buffer{}
	var handler slog.Handler
	handler = slog.NewTextHandler(buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	handler = WithRecordID("traceId", handler)
	slog.SetDefault(slog.New(handler))

	ctx1 := context.Background()
	ctx1 = SetID(ctx1, "1234")

	slog.InfoContext(ctx1, "test")
	if buf.String() == "" {
		t.Error("buf.String() == ")
	}
	if !strings.Contains(buf.String(), `traceId=1234`) {
		t.Logf("buf.String() == %s", buf.String())
		t.Error(`!strings.Contains(buf.String(), "traceId=1234")`)
	}
}
