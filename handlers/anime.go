package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/suppayami/pettanko.go/commons"
	"github.com/suppayami/pettanko.go/repositories"
)

// AnimeHandler responses to /anime/ resource
type AnimeHandler struct {
	animeRepo repositories.AnimeRepository
}

// ServeHTTP implement
func (handler *AnimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handler.getMethod(w, r)
		return
	}
	commons.MethodNotAllowedHandler(w, r)
}

func (handler *AnimeHandler) getMethod(w http.ResponseWriter, r *http.Request) {
	// resource root
	routeParams := strings.Split(r.URL.Path, "/")[2:]
	if routeParams[0] == "" {
		handler.allAnime(w, r)
		return
	}

	// resource at id
	id, err := strconv.Atoi(routeParams[0])
	if err != nil {
		commons.NotFoundHandler(w, r)
		return
	}
	handler.anime(w, r, id)
}

func (handler *AnimeHandler) allAnime(w http.ResponseWriter, r *http.Request) {
	animeList := handler.animeRepo.AllAnime()
	json.NewEncoder(w).Encode(animeList)
}

func (handler *AnimeHandler) anime(w http.ResponseWriter, r *http.Request, id int) {
	anime := handler.animeRepo.Anime(id)
	if anime == nil {
		commons.NotFoundHandler(w, r)
		return
	}
	json.NewEncoder(w).Encode(anime)
}
