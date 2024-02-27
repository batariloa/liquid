package server

import (
	"log"

	"github.com/batariloa/search-service/internal/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	sh     *handler.SearchHandler
}

func NewServer(searchHandler *handler.SearchHandler) *Server {
	router := gin.Default()

	s := &Server{
		router: router,
		sh:     searchHandler,
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	s.router.GET("/search/songs/:query", s.sh.HandleSearchSongByTitleAndArtist)
}

func (s *Server) Run(addr string) error {
	log.Printf("Service starting on %s", addr)
	return s.router.Run(addr)
}
