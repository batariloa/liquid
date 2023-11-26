package artist

import (
	"StorageService/test"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func TestCreateartist(t *testing.T) {

	db, err := test.SetUpDbContainer(t)

	defer db.Close()

	if err != nil {
		t.Fatal(err)
	}

	repository := NewPqlRepository(db)
	service := NewArtistService(repository)
	handler := NewArtistHandler(service)

	expectedName := "Some Name"

	input := map[string]interface{}{
		"name": expectedName,
	}

	requestBody, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	router := SetUpRouter()
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

	artist, err := repository.GetById(1)
	if err != nil {
		t.Fatalf("Failed getting artist %s", err)
	}

	assert.Equal(t, artist.Name, expectedName)
}
