package commons

import (
	"encoding/json"
	"net/http"

	"github.com/suppayami/pettanko.go/models"
)

// ForbiddenHandler handles Forbidden (403) error
func ForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "403 forbidden"}
	w.WriteHeader(403)
	json.NewEncoder(w).Encode(msg)
}

// NotFoundHandler handles Not Found (404) error
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "404 not found"}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(msg)
}

// MethodNotAllowedHandler handles Method Not Allowed (405) error
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "405 method not allowed"}
	w.WriteHeader(405)
	json.NewEncoder(w).Encode(msg)
}
