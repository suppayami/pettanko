package handlers

import (
	"log"
	"net/http"

	"github.com/suppayami/pettanko/middleware"
	"github.com/suppayami/pettanko/service"
)

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
}

// SetupRoute setups handles for each route and is called in main
func SetupRoute() {
	mux.Handle("/", middleware.ResponseJSON(&IndexHandler{}))
	mux.Handle("/auth/", middleware.ResponseJSON(&AuthHandler{}))
	mux.Handle("/ping/", middleware.ResponseJSON(&PingHandler{}))
	mux.Handle("/blocked/", middleware.ResponseJSON(
		middleware.BlockRoute(&PingHandler{}),
	))
	mux.Handle("/anime/", middleware.ResponseJSON(
		&AnimeHandler{animeRepo: service.AnimeRepository()},
	))
}

// Serve listens to given address and handles incoming requests
func Serve(addr string) {
	log.Printf("Listening to %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
