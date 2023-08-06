package helper

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error converting response to JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
