package jibe

import (
	"log/slog"
	"net/http"
	"time"
)

// struct for holding response details
type responseData struct {
	status int
	size   int
}

type responseWrapper struct {
	http.ResponseWriter
	responseData *responseData
}

func (r *responseWrapper) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size = size
	return size, err
}

func (r *responseWrapper) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func WithLogging(level slog.Level, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rd := &responseData{}
		lw := &responseWrapper{ResponseWriter: w, responseData: rd}
		next.ServeHTTP(lw, r)
		duration := time.Since(start)
		slog.Log(r.Context(), level, "http request", "duration", duration, "method", r.Method, "size", rd.size, "status", rd.status, "uri", r.RequestURI)
	})
}
