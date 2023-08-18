package service

import (
	"StorageService/internal/datastruct"
	"StorageService/internal/util"
	"StorageService/internal/util/apierror"
	"fmt"
	"mime/multipart"
	"path/filepath"
)

type UploadService struct {
	KafkaService *KafkaService
}

func NewUploadService(ks *KafkaService) *UploadService {
	return &UploadService{
		KafkaService: ks,
	}
}

func (u *UploadService) UploadSong(file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {

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

func (u *UploadService) GenerateAndPublishSongUploadEvent(artistID int, title, artistName string) error {

	uploadEvent := datastruct.UploadKafkaEvent{
		ArtistName: artistName,
		SongName:   title,
		SongID:     artistID,
	}

	err := u.KafkaService.PublishUploadSongEvent(uploadEvent)
	if err != nil {
		return err
	}

	return nil
}
