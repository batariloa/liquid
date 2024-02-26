package handlers

import (
	"StorageService/internal/apierror"
	"StorageService/internal/data"
	"StorageService/internal/service"
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
	name := r.FormValue("name")
	if name == "" {
		return "", apierror.NewBadRequestError("Artist name not provided")
	}

	return name, nil
}
