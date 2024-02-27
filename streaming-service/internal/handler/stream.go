package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	service "github.com/batariloa/StreamingService/internal/service/fetcher"
	"github.com/gin-gonic/gin"
)

type StreamHandler struct {
	fetcherService service.SongFetcher
}

func New(fs service.SongFetcher) *StreamHandler {
	return &StreamHandler{
		fetcherService: fs,
	}
}

// Stream Song godoc
// @Summary Stream a song to user
// @Description Stream a song to the user by providing the song ID
// @ID streamSongToUser
// @Produce octet-stream
// @Param songId path integer true "Song ID to stream"
// @Success 200 {file} application/octet-stream
// @Failure 400 {string} string "Bad Request: Invalid ID"
// @Failure 500 {string} string "Internal Server Error: File download failed"
// @Router /v1/stream/{songId} [get]
func (h *StreamHandler) StreamSong(c *gin.Context) {

	idStr := c.Param("songId")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
	}

	response, err := h.fetcherService.Fetch(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "File download failed")
		return
	}
	defer response.Body.Close()

	c.Header("Content-Type", "audio/mpeg")

	_, err = io.Copy(c.Writer, response.Body)
	if err != nil {
		fmt.Println("Error while streaming:", err)
	}
}
