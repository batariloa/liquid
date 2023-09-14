package handler

import (
	"github.com/batariloa/search-service/internal/search"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	searchService search.SearchService
}

func NewSearchHandler(searchService *search.SearchService) *SearchHandler {

	return &SearchHandler{
		searchService: *searchService,
	}
}

func (h *SearchHandler) HandleSearchByTitleAndArtist(c *gin.Context) {
	query := c.Param("query")
	songs, err := h.searchService.SearchSongsByTitleOrArtist(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"songs": songs})
}
