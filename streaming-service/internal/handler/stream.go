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

// StreamSong godoc
// @Summary Stream a song to user
// @Description Stream a song to the user by providing the song ID
// @Tags StreamHandler
// @ID streamSongToUser
// @Produce octet-stream
// @Param songId path int true "Song ID to stream"
// @Success 200 {file} application/octet-stream
// @Failure 400 {string} string "Bad Request: Invalid ID"
// @Failure 500 {string} string "Internal Server Error: File download failed"
// @Router /v1/stream/{songId} [get]
func (h StreamHandler) StreamSong(c *gin.Context) {

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
	c.Status(http.StatusOK)

	buf := make([]byte, 1024*10) // 10KB buffer size
	for {
		n, err := response.Body.Read(buf)
		if n > 0 {
			if _, writeErr := c.Writer.Write(buf[:n]); writeErr != nil {
				fmt.Println("Error while streaming:", writeErr)
				return
			}
			c.Writer.Flush() // Ensure that the data is sent to the client
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error while reading:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error: Streaming failed")
			return
		}
	}
}
