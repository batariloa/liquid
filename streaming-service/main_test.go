package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batariloa/StreamingService/internal/handler"
	"github.com/batariloa/StreamingService/internal/test"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSongHandler_getSongById(t *testing.T) {

	fetchService := &test.MockFetcherService{}
	handler := handler.New(fetchService)

	r := SetUpRouter()
	r.GET("/:songId", handler.StreamSong)
	req, _ := http.NewRequest("GET", "/3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
