package service

import (
	"net/http"
)

type SongFetcher interface {
	Fetch(id int) (*http.Response, error)
}
