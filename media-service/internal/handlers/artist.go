package handlers

import (
	"StorageService/internal/apierror"
	"StorageService/internal/data"
	"StorageService/internal/service"
	"encoding/json"
	"net/http"
)

func (*Handler) HandleCreateArtist(w http.ResponseWriter, r *http.Request) {

	name, err := getArtistName(r)
	if err != nil {
		http.Error(w, "Please provide all required fields", http.StatusBadRequest)
		return
	}

	artistInput := data.Artist{Name: name}

	res, err := service.SaveArtist(&artistInput)
	if err != nil {
		http.Error(w, "Error creating artist", http.StatusInternalServerError)
		return
	}

	WriteJSONResponse(w, res, http.StatusCreated)
}

func getArtistName(r *http.Request) (string, error) {
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return "", apierror.NewBadRequestError("Error decoding request body")
	}

	name, ok := requestBody["name"].(string)
	if !ok {
		return "", apierror.NewBadRequestError("Artist name not provided or invalid")
	}

	return name, nil
}
