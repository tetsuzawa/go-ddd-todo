package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("endpoint=%s, method=%s duration=%d",
			r.URL.Path,
			r.Method,
			time.Since(start),
		)
	}

	return http.HandlerFunc(fn)
}
