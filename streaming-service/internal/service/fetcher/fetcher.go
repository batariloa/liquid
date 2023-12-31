package service

import (
	"fmt"
	"log"
	"net/http"
)

var baseURL = "http://media-service:3000/songs/%d/download"

type SongFetcherService struct {
}

func NewFetcherService() *SongFetcherService {
	return &SongFetcherService{}
}

func (*SongFetcherService) Fetch(id int) (*http.Response, error) {

	downloadURL := fmt.Sprintf(baseURL, id)
	log.Println("Download url", downloadURL)

	response, err := http.Get(downloadURL)
	if err != nil {
		log.Println("Error getting song", err)
		return nil, fmt.Errorf("Error whle making the GET request: %s\n", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: GET request returned status code %d\n", response.StatusCode)
	}

	return response, nil
}
