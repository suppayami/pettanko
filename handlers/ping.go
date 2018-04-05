package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/suppayami/pettanko.go/models"
)

// PingHandler responses all requests with pong
type PingHandler struct{}

// ServeHTTP implement
func (handler *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "pong"}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msg)
}
