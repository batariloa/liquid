package server

import (
	"github.com/batariloa/search-service/internal/search"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	searchSvc *search.SearchService
}

func NewServer(searchSvc *search.SearchService) *Server {
	router := gin.Default()

	s := &Server{
		router:    router,
		searchSvc: searchSvc,
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	s.router.GET("/search/:query", s.searchByTitleAndArtist)
}

func (s *Server) searchByTitleAndArtist(c *gin.Context) {
	query := c.Param("query")
	songs, err := s.searchSvc.SearchSongsByTitleOrArtist(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"songs": songs})
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}
