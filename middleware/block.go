package middleware

import (
	"net/http"

	"github.com/suppayami/pettanko.go/commons"
)

// BlockRoute blocks all requests to given handler
func BlockRoute(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commons.ForbiddenHandler(w, r)
	})
}
