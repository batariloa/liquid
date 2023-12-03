// server/server.go
package server

import (
	"StorageService/internal/handlers"
	"StorageService/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	EventProducer service.EventPublisher
}

func NewAPI(eventProducer service.EventPublisher) *API {
	return &API{
		EventProducer: eventProducer,
	}
}

func (s *API) RegisterRoutes() *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/artists", handlers.HandleCreateArtist).Methods("POST")
	router.HandleFunc("/songs", handlers.HandleGetSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", handlers.HandleGetSongByID).Methods("GET")
	router.HandleFunc("/songs", handlers.HandleUploadSong).Methods("POST")
	router.HandleFunc("/songs/{id}/download", handlers.HandleDownloadSong).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	return server
}
