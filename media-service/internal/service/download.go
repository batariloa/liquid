package service

import (
	"StorageService/internal/apierror"
	"StorageService/internal/data"
	"os"
)

func DownloadSongById(songId int) (*os.File, error) {

	songData, err := data.GetSongById(songId)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(songData.FilePath)
	if err != nil {
		return nil, apierror.NewNotFoundError("Song file could not be found.")
	}

	return file, nil
}
