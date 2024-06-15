package main

import (
	"io/ioutil"
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

	assert.Equal(t, "audio/mpeg", w.Header().Get("Content-Type"))

	body, _ := ioutil.ReadAll(w.Body)
	expectedContent := "Mock MP3 file content goes here"
	assert.Equal(t, expectedContent, string(body))
}
