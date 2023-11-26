package artist

import (
	"StorageService/internal/handlers/helper"
	"log"
	"net/http"
)

type ArtistHandler struct {
	ArtistService ArtistService
}

func NewArtistHandler(artistService ArtistService) *ArtistHandler {
	return &ArtistHandler{
		ArtistService: artistService,
	}
}

func (h *ArtistHandler) HandleCreateArtist(w http.ResponseWriter, r *http.Request) {

	log.Print("Starting")
	name, err := helper.GetArtistName(r)
	if err != nil {
		http.Error(w, "Please provide all required fields", http.StatusBadRequest)
		return
	}

	log.Print("Creating artist object")

	artistInput := Artist{Name: name}

	res, err := h.ArtistService.Save(&artistInput)
	if err != nil {
		http.Error(w, "Error creating artist", http.StatusInternalServerError)
		return
	}

	helper.WriteJSONResponse(w, res, http.StatusCreated)
}
