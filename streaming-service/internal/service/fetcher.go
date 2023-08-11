package service

import (
	"fmt"
	"io"
	"net/http"
)

var downloadUrl = "http://localhost:3000/songs/%s/download"

type SongFetcherService struct {
}

func NewFetcherService() *SongFetcherService {
	return &SongFetcherService{}
}

func (*SongFetcherService) FetchSongFileResponseById(id int) (*http.Response, error) {

	downloadUrl := fmt.Sprintf(downloadUrl, id)

	response, err := http.Get(downloadUrl)
	if err != nil {
		return nil, fmt.Errorf("Error whle making the GET request: %s\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: GET request returned status code %d\n", response.StatusCode)
	}

	return response, nil
}

func (*SongFetcherService) WriteSongContentToResponse(w io.Writer, body io.Reader) {
	_, err := io.Copy(w, body)
	if err != nil {
		fmt.Println("Error while streaming:", err)
	}
}
