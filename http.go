package jibe

import (
	"net/http"
)

// WithRequestID adds a unique ID to the HTTP request context.
func WithRequestID(makeID IDProducer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := SetID(r.Context(), makeID())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
