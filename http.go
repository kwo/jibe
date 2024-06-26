package jibe

import (
	"net/http"
)

const HeaderRequestID = "X-Request-ID"

type IDProducer func() string

// WithRequestID adds a unique ID to the HTTP request context.
func WithRequestID(makeID IDProducer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := SetID(r.Context(), makeID())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// WithResponseHeader adds the ID from the request context to the HTTP response.
func WithResponseHeader(headerName string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id := GetID(r.Context()); id != "" {
			w.Header().Set(headerName, id)
		}
		next.ServeHTTP(w, r)
	})
}
