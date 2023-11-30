package service

import (
	"StorageService/internal/data"
	"StorageService/internal/util/apierror"
	"fmt"
)

func Save(song *data.SongData) (*data.SongData, error) {
	res, err := data.SaveSong(song)
	if err != nil {
		fmt.Println("Trouble saving song data", err)
		return nil, apierror.NewInternalServerError("Could not save song data.")
	}

	return res, nil
}

func GetSongById(songId int) (*data.SongData, error) {

	song, err := data.GetSongById(songId)
	if err != nil {
		return nil, err
	}

	if song == nil {
		return nil, apierror.NewNotFoundError("Song not found")
	}

	return song, nil
}
