package handler

import (
	"net/http"
	"strconv"

	"github.com/batariloa/StreamingService/internal/service"
	"github.com/gin-gonic/gin"
)

type StreamHandler struct {
	fetcherService service.SongFetcherService
}

func New(fs *service.SongFetcherService) *StreamHandler {
	return &StreamHandler{
		fetcherService: *fs,
	}
}

func (h *StreamHandler) StreamFileToUserHandler(c *gin.Context) {

	idStr := c.Param("songId")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
	}

	response, err := h.fetcherService.FetchSongFileResponseById(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "File download failed")
		return
	}
	defer response.Body.Close()

	h.fetcherService.WriteSongContentToResponse(c.Writer, response.Body)
}
