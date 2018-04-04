package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/suppayami/pettanko/models"
)

// ForbiddenHandler handles Forbidden (403) error
func ForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "403 forbidden"}
	w.WriteHeader(403)
	json.NewEncoder(w).Encode(msg)
}
