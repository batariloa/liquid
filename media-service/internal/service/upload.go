package service

import (
	"StorageService/internal/apierror"
	"StorageService/internal/types"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UploadSong(file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	fileName := fileHeader.Filename
	fileExt := filepath.Ext(fileName)

	if fileExt != ".mp3" {
		return "", apierror.NewBadRequestError("Not an mp3 file")
	}

	uniqueFileName := generateRandomFileName(fileExt)

	filePath, err := saveFile(*file, uniqueFileName)
	if err != nil {
		fmt.Print(err)
		return "", apierror.NewInternalServerError("Error uploading the file.")
	}

	return filePath, nil
}

func GenerateAndPublishSongUploadEvent(event types.UploadSongEvent, publisher EventPublisher) error {

	err := publisher.PublishUploadSongEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func saveFile(file io.Reader, fileName string) (string, error) {
	destFolder := "files/tracks" // Remove the dot (.) from the folder path

	// Create the destination folder and any necessary parent directories
	err := os.MkdirAll(destFolder, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create destination folder: %w", err)
	}

	// Create the full file path by joining the destination folder and the filename
	filePath := filepath.Join(destFolder, fileName)
	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// Copy the file content from the reader to the created file
	_, err = io.Copy(f, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	return filePath, nil
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generateRandomFileName(fileExt string) string {

	uniqueID := generateRandomString(10)
	timeStamp := time.Now().Format("20060102150405")
	uniqueFileName := fmt.Sprintf("%s_%s%s", timeStamp, uniqueID, fileExt)

	return uniqueFileName
}
