package server

import (
	"StorageService/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	storageHandler  *handlers.SongHandler
	artistHandler   *handlers.ArtistHandler
	downloadHandler *handlers.DownloadHandler
}

func NewServer(sh *handlers.SongHandler, ah *handlers.ArtistHandler, dh *handlers.DownloadHandler) *Server {
	return &Server{
		storageHandler:  sh,
		artistHandler:   ah,
		downloadHandler: dh,
	}
}

func (s *Server) SetupServer() *http.Server {

	router := mux.NewRouter()

	router.HandleFunc("/artists", s.artistHandler.HandleCreateArtist).Methods("POST")
	router.HandleFunc("/songs", s.storageHandler.GetSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", s.storageHandler.GetSongByID).Methods("GET")
	router.HandleFunc("/songs", s.storageHandler.UploadSong).Methods("POST")
	router.HandleFunc("/songs/{id}/download", s.downloadHandler.HandleDownloadSong).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	return server
}
