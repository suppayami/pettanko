package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/suppayami/pettanko/models"
)

// NotFoundHandler handles Not Found (404) error
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	msg := &models.SimpleMessage{Message: "404 not found"}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(msg)
}
