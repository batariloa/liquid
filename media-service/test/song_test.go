package test

import (
	"StorageService/internal/handlers"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockKafkaService struct {
	mock *mock.Mock
}

func TestUploadSong(t *testing.T) {
	tstDb, err := setUpDbContainer(t)
	if err != nil {
		t.Fatal(err)
	}
	defer tstDb.Close()

	router := setUpRouter()
	router.HandleFunc("/songs", handlers.HandleUploadSong).Methods("POST")

	mp3FilePath := "./resources/example_song_file.mp3"
	mp3File, err := os.Open(mp3FilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer mp3File.Close()

	// Create a buffer to hold the multipart content
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Add the MP3 file to the multipart content with the desired field name "song_file"
	filePart, err := writer.CreateFormFile("song-file", filepath.Base(mp3FilePath))
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(filePart, mp3File)
	if err != nil {
		t.Fatal(err)
	}

	// Add additional form fields
	writer.WriteField("title", "Example Song Title")
	writer.WriteField("artistId", "5")

	// Close the writer to finalize the multipart content
	writer.Close()

	// Create the POST request with the multipart content
	req, err := http.NewRequest("POST", "/songs", &body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("Response Status Code: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)
}
