package middleware

import (
	"net/http"
)

// GetMethod allows get method
func GetMethod(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			MethodNotAllowedHandler(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// PostMethod allows post method
func PostMethod(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "Post" {
			MethodNotAllowedHandler(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
