package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/suppayami/pettanko.go/commons"
	"github.com/suppayami/pettanko.go/models"
)

// IndexHandler responses to root path ("/")
type IndexHandler struct{}

// ServeHTTP implement
func (handler *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		commons.NotFoundHandler(w, r)
		return
	}
	msg := &models.SimpleMessage{Message: "hello world from index"}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(msg)
}
