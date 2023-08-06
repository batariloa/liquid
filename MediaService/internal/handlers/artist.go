package handlers

import (
	"StorageService/internal/handlers/helper"
	"StorageService/internal/repository/artist"
	"StorageService/internal/service"
	"net/http"
)

type ArtistHandler struct {
	ArtistService *service.ArtistService
}

func NewArtistHandler(artistService *service.ArtistService) *ArtistHandler {
	return &ArtistHandler{
		ArtistService: artistService,
	}
}

func (h *ArtistHandler) HandleCreateArtist(w http.ResponseWriter, r *http.Request) {

	name, err := helper.GetArtistName(r)
	if err != nil {
		http.Error(w, "Please provide all required fields", http.StatusBadRequest)
	}

	artistInput := artist.Artist{Name: name}

	res, err := h.ArtistService.Save(&artistInput)
	if err != nil {
		http.Error(w, "Error creating artist", http.StatusInternalServerError)
	}

	helper.WriteJSONResponse(w, res, http.StatusCreated)
}
