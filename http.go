package jibe

import "net/http"

// Middleware adds a unique ID to the HTTP request context.
func Middleware(makeID IDProducer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := SetID(r.Context(), makeID())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
