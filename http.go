package jibe

import (
	"net/http"
)

const HeaderRequestID = "X-Request-ID"

type IDProducer func() string

// WithRequestID adds a unique ID to the HTTP request context.
func WithHTTPHandler(makeID IDProducer, headerName string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := makeID()
		ctx := SetID(r.Context(), id)
		w.Header().Set(headerName, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
