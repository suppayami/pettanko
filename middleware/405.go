package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/suppayami/pettanko/models"
)

// MethodNotAllowedHandler handles Method Not Allowed (405) error
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "405 method not allowed"}
	w.WriteHeader(405)
	json.NewEncoder(w).Encode(msg)
}
