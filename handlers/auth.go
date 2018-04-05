package handlers

import (
	"net/http"
	"strings"

	"github.com/suppayami/pettanko/service"

	"github.com/suppayami/pettanko/commons"
)

// AuthHandler responses to authentication requests
type AuthHandler struct {
	authService *service.Authentication
}

// ServeHTTP implement
func (handler *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handler.postMethod(w, r)
		return
	}
	commons.MethodNotAllowedHandler(w, r)
}

func (handler *AuthHandler) postMethod(w http.ResponseWriter, r *http.Request) {
	routeParams := strings.Split(r.URL.Path, "/")[2:]

	if routeParams[0] == "login" {
		handler.login(w, r)
		return
	}

	if routeParams[0] == "logout" {
		handler.logout(w, r)
		return
	}

	commons.NotFoundHandler(w, r)
}

func (handler *AuthHandler) login(w http.ResponseWriter, r *http.Request) {

}

func (handler *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {

}
