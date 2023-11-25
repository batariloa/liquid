package upload

import (
	"StorageService/internal/datastruct"
	"StorageService/internal/kafka"
	"StorageService/internal/util"
	"StorageService/internal/util/apierror"
	"fmt"
	"mime/multipart"
	"path/filepath"
)

type UploadService struct {
	KafkaService *kafka.KafkaService
}

func NewUploadService(ks *kafka.KafkaService) *UploadService {
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

func (u *UploadService) GenerateAndPublishSongUploadEvent(songId int, title, artistName string) error {

	uploadEvent := datastruct.UploadKafkaEvent{
		ArtistName: artistName,
		Title:      title,
		SongID:     songId,
	}

	err := u.KafkaService.PublishUploadSongEvent(uploadEvent)
	if err != nil {
		return err
	}

	return nil
}
