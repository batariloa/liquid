package helper

import (
	"StorageService/internal/apierror"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetArtistID(r *http.Request) (int, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MBs
	if err != nil {
		return 0, apierror.NewBadRequestError("Error parsing multipart form")
	}

	artistID, err := strconv.Atoi(r.FormValue("artistId"))
	if err != nil {
		return 0, apierror.NewBadRequestError("Artist ID not provided or invalid")
	}

	return artistID, nil
}

func GetSongTitle(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MBs
	if err != nil {
		return "", apierror.NewBadRequestError("Error parsing multipart form")
	}

	title := r.FormValue("title")
	if title == "" {
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
