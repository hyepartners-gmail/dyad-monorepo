package middleware

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func WithCommonMiddleware(handler http.Handler) http.Handler {
	return cors.AllowAll().Handler(Logging(handler))
}
