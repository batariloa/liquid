package helper

import (
	"StorageService/internal/apierror"
	"encoding/json"
	"net/http"
)

func GetArtistID(r *http.Request) (int, error) {
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return 0, apierror.NewBadRequestError("Error decoding request body")
	}

	artistID, ok := requestBody["artistId"].(float64)
	if !ok {
		return 0, apierror.NewBadRequestError("Artist ID not provided or invalid")
	}

	return int(artistID), nil
}

func GetSongTitle(r *http.Request) (string, error) {
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return "", apierror.NewBadRequestError("Error decoding request body")
	}

	title, ok := requestBody["title"].(string)
	if !ok {
		return "", apierror.NewBadRequestError("Song title not provided or invalid")
	}

	return title, nil
}

func GetArtistName(r *http.Request) (string, error) {
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

func GetSongId(r *http.Request) (string, error) {
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return "", apierror.NewBadRequestError("Error decoding request body")
	}

	id, ok := requestBody["songId"].(string)
	if !ok {
		return "", apierror.NewBadRequestError("Song ID not provided or invalid")
	}

	return id, nil
}
