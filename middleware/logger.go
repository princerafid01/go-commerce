package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("Ami Middleware prothome run hobo")

		next.ServeHTTP(w, r)

		diff := time.Since(start)
		log.Println("Ami Middleware ses e  run hobo")

		log.Println(r.Method, r.URL.Path, diff)
	})
}
