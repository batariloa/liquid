package service

import (
	"StorageService/internal/types"
	"StorageService/internal/util"
	"StorageService/internal/util/apierror"
	"fmt"
	"mime/multipart"
	"path/filepath"
)

func UploadSong(file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	fileName := fileHeader.Filename
	fileExt := filepath.Ext(fileName)

	if fileExt != ".mp3" {
		return "", apierror.NewBadRequestError("Not an mp3 file")
	}

	uniqueFileName := util.GenerateRandomFileName(fileExt)

	filePath, err := util.SaveFile(*file, uniqueFileName)
	if err != nil {
		fmt.Print(err)
		return "", apierror.NewInternalServerError("Error uploading the file.")
	}

	return filePath, nil
}

func GenerateAndPublishSongUploadEvent(songId int, title, artistName string) error {

	uploadEvent := types.UploadSongEvent{
		ArtistName: artistName,
		Title:      title,
		SongID:     songId,
	}

	err := PublishUploadSongEvent(uploadEvent)
	if err != nil {
		return err
	}

	return nil
}
