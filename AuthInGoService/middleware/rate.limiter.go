package middleware

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var limiter=rate.NewLimiter(rate.Every(1*time.Second),5) //5 request per second

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}