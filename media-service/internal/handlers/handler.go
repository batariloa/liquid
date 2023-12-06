package handlers

import (
	"StorageService/internal/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	EventPublisher service.EventPublisher
}

func NewHandler(eventPublisher service.EventPublisher /*, other parameters */) *Handler {
	return &Handler{
		EventPublisher: eventPublisher,
	}
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error converting response to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
