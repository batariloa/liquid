package search

import (
	"fmt"

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
	queryString := "title:" + query + " OR artist:" + query
	searchQuery := bleve.NewQueryStringQuery(queryString)
	searchRequest := bleve.NewSearchRequest(searchQuery)

	searchResults, err := s.index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	// Collect the matching songs
	var songs []model.Song
	for _, hit := range searchResults.Hits {
		songs = append(songs, model.Song{
			ID:     hit.ID,
			Title:  hit.Fields["title"].(string),
			Artist: hit.Fields["artist"].(string),
		})
	}

	return songs, nil
}

func (s *SearchService) IndexSong(song model.Song) error {
	doc := map[string]interface{}{
		"id":     song.ID,
		"title":  song.Title,
		"artist": song.Artist,
	}
	return s.index.Index(fmt.Sprintf("%s - %s", song.Title, song.Artist), doc)
}
