package middleware

import (
	"log"
	"net/http"
)

// LogrequestHandler will log the HTTP requests
func LogrequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.URL)
		next.ServeHTTP(w, r)
	})
}
