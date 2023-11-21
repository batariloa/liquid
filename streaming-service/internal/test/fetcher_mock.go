package test

import (
	"io"
	"net/http"
	"strings"
)

type MockFetcherService struct {
}

func (m *MockFetcherService) Fetch(id int) (*http.Response, error) {

	mockMP3Content := "Mock MP3 file content goes here"
	mockResponse := &http.Response{
		Body: io.NopCloser(strings.NewReader(mockMP3Content)),
		Header: map[string][]string{
			"Content-Type": {"audio/mpeg"},
		},
	}

	return mockResponse, nil
}
