package test

import (
	"StorageService/internal/data"
	"StorageService/internal/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateArtist(t *testing.T) {

	tstDb, err := setUpDbContainer(t)
	defer tstDb.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedName := "Some Name"

	input := map[string]interface{}{
		"name": expectedName,
	}

	requestBody, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	mockEventPublsher := new(MockEventPublisher)
	handler := handlers.NewHandler(mockEventPublsher)

	router := setUpRouter()
	router.HandleFunc("/artists", handler.HandleCreateArtist).Methods("POST")

	req, err := http.NewRequest("POST", "/artists", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("Response Status Code: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)

	artist, err := data.GetArtistById(1)
	if err != nil {
		t.Fatalf("Failed getting artist %s", err)
	}

	assert.Equal(t, artist.Name, expectedName)
}
