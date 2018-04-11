package middleware

import (
	"net/http"

	"github.com/suppayami/pettanko/commons"

	"github.com/suppayami/authgo"
	"github.com/suppayami/pettanko/service"
)

var (
	localAuth func(http.Handler) http.Handler
)

func init() {
	localAuth = authgo.MakeMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			commons.UnauthorizedHandler(w, r)
		}),
		service.LocalStrategy,
	)
}

// LocalAuth check authentication with local strategy from authgo
func LocalAuth(handler http.Handler) http.Handler {
	return localAuth(handler)
}
