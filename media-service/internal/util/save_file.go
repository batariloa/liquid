package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SaveFile(file io.Reader, fileName string) (string, error) {
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
