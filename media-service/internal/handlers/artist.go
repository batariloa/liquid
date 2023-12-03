package handlers

import (
	"StorageService/internal/data"
	"StorageService/internal/handlers/helper"
	"StorageService/internal/service"
	"net/http"
)

func HandleCreateArtist(w http.ResponseWriter, r *http.Request) {

	name, err := helper.GetArtistName(r)
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

	helper.WriteJSONResponse(w, res, http.StatusCreated)
}
