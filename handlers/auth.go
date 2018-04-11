package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/suppayami/pettanko/middleware"
	"github.com/suppayami/pettanko/models"

	"github.com/suppayami/pettanko/commons"
)

// AuthHandler responses to authentication requests
type AuthHandler struct {
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
		middleware.LocalAuth(http.HandlerFunc(handler.login)).ServeHTTP(w, r)
		return
	}

	commons.NotFoundHandler(w, r)
}

func (handler *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "login successfully"}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msg)
}
