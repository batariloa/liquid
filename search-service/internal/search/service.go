package search

import (
	"log"

	"github.com/batariloa/search-service/internal/model"
	"github.com/blevesearch/bleve/v2"
)

type SearchService struct {
	index bleve.Index
}

func NewSearchService(index bleve.Index) *SearchService {
	return &SearchService{index: index}
}

func (s *SearchService) SearchSongsByTitleOrArtist(query string) ([]model.Song, error) {
	log.Printf("Searching for title or artist: %s", query)

	queryString := "title:" + query + " OR artist:" + query
	searchQuery := bleve.NewQueryStringQuery(queryString)
	searchRequest := bleve.NewSearchRequest(searchQuery)
	searchRequest.Fields = []string{"title", "artist"}
	searchRequest.Size = 10

	searchResults, err := s.index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	// Collect the matching songs
	var songs []model.Song
	for _, hit := range searchResults.Hits {
		song := model.Song{
			ID: hit.ID,
		}
		if title, ok := hit.Fields["title"].(string); ok {
			song.Title = title
		}
		if artist, ok := hit.Fields["artist"].(string); ok {
			song.ArtistName = artist
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (s *SearchService) IndexSong(song model.Song) error {
	doc := map[string]interface{}{
		"id":         song.ID,
		"title":      song.Title,
		"artist":     song.ArtistName,
		"uploadedBy": song.UploadedBy,
	}
	return s.index.Index(song.ID, doc)
}
