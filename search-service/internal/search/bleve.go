package search

import (
	"log"

	"github.com/blevesearch/bleve/v2"
)

func InitializeBleveIndex() bleve.Index {

	mapping := bleve.NewIndexMapping()

	var err error
	index, err := bleve.New("songs.bleve", mapping)
	if err != nil {
		log.Fatal("Error creating Bleve index:", err)
	}

	return index
}
