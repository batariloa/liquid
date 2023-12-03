package service

import (
	"StorageService/internal/apierror"
	"StorageService/internal/data"
)

func GetArtistById(artistId int) (*data.Artist, error) {

	result, err := data.GetArtistById(artistId)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, apierror.NotFoundError{Message: "Artist not found."}
	}

	return result, nil
}

func SaveArtist(artist *data.Artist) (*data.Artist, error) {

	res, err := data.SaveArtist(artist)
	if err != nil {
		return nil, err
	}
	return res, nil
}
