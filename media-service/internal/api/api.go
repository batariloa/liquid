package api

import (
	"StorageService/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	handler *handlers.Handler
}

func NewAPI(h *handlers.Handler) *API {
	return &API{
		handler: h,
	}
}

func (s *API) RegisterRoutes() *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/artists", s.handler.HandleCreateArtist).Methods("POST")
	router.HandleFunc("/songs/{id}", s.handler.HandleGetSongByID).Methods("GET")
	router.HandleFunc("/songs", s.handler.HandleUploadSong).Methods("POST")
	router.HandleFunc("/songs/{id}/download", s.handler.HandleDownloadSong).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	return server
}
