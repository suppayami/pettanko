package middleware

import (
	"net/http"
)

// ResponseJSON makes handler responses with Content-Type application/json
func ResponseJSON(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		handler.ServeHTTP(w, r)
	})
}
