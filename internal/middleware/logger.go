package middleware

import (
	"log"
	"net/http"
	"time"
)

// Custom response writer to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Override WriteHeader to capture status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		// Wrap original ResponseWriter
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // default status
		}

		// Log incoming request
		log.Printf("[REQUEST] %s %s", r.Method, r.URL.Path)

		// Call next handler
		next.ServeHTTP(rw, r)

		// Log response details
		duration := time.Since(start).Microseconds()

		log.Printf(
			"[RESPONSE] %d %s %s took %dµs",
			rw.statusCode,
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}
