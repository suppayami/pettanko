package middleware

import (
	"net/http"

	"github.com/suppayami/pettanko/commons"
	"github.com/suppayami/pettanko/service"
)

// AuthRoute check authentication status
func AuthRoute(handler http.Handler, auth service.Authentication) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !auth.LoggedIn() {
			commons.UnauthorizedHandler(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
