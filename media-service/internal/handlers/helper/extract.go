package helper

import (
	"StorageService/internal/util/apierror"
	"net/http"
	"strconv"
)

func GetArtistID(r *http.Request) (int, error) {
	artistIdStr := r.FormValue("artistId")
	if artistIdStr == "" {
		return 0, apierror.NewBadRequestError("Artist ID not provided")
	}
	return strconv.Atoi(artistIdStr)
}

func GetSongTitle(r *http.Request) (string, error) {
	title := r.FormValue("title")
	if title == "" {
		return "", apierror.NewBadRequestError("Song title not provided")
	}
	return title, nil
}

func GetArtistName(r *http.Request) (string, error) {
	name := r.FormValue("name")
	if name == "" {
		return "", apierror.NewBadRequestError("Artist name not provided")
	}
	return name, nil
}

func GetSongId(r *http.Request) (string, error) {

	id := r.FormValue("songId")
	if id == "" {
		return "", apierror.NewBadRequestError("Song Id not provided")
	}

	return id, nil
}
