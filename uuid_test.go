package jibe

import "testing"

func TestUuid(t *testing.T) {
	id := NewID()
	if id == "" {
		t.Error("id == ")
	}
	if len(id) != 36 {
		t.Error("len(id) != 36")
	}
}
