package jibe

import "github.com/google/uuid"

func NewID() string {
	id, err := uuid.NewV7()
	if err != nil {
		return "no-id"
	}
	return id.String()
}
