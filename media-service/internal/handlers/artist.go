package handlers

import (
	"StorageService/internal/artist"
	"StorageService/internal/handlers/helper"
	"net/http"
)

type ArtistHandler struct {
	ArtistService artist.ArtistService
}

func NewArtistHandler(artistService artist.ArtistService) *ArtistHandler {
	return &ArtistHandler{
		ArtistService: artistService,
	}
}

func (h *ArtistHandler) HandleCreateArtist(w http.ResponseWriter, r *http.Request) {

	name, err := helper.GetArtistName(r)
	if err != nil {
		http.Error(w, "Please provide all required fields", http.StatusBadRequest)
		return
	}

	artistInput := artist.Artist{Name: name}

	res, err := h.ArtistService.Save(&artistInput)
	if err != nil {
		http.Error(w, "Error creating artist", http.StatusInternalServerError)
		return
	}

	helper.WriteJSONResponse(w, res, http.StatusCreated)
}
