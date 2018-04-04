package middleware

import (
	"net/http"
)

// BlockRoute blocks all requests to given handler
func BlockRoute(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ForbiddenHandler(w, r)
	})
}
