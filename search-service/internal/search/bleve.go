package search

import (
	"log"
	"os"

	"github.com/blevesearch/bleve/v2"
)

const indexDirectory = "songs.bleve"

func InitializeBleveIndex() (bleve.Index, error) {
	indexExists, err := indexDirectoryExists(indexDirectory)
	if err != nil {
		return nil, err
	}

	var index bleve.Index

	if indexExists {
		index, err = bleve.Open(indexDirectory)
		if err != nil {
			log.Fatal("Error opening existing Bleve index:", err)
			return nil, err
		}
	} else {

		mapping := bleve.NewIndexMapping()

		index, err = bleve.New(indexDirectory, mapping)
		if err != nil {
			log.Fatal("Error creating Bleve index:", err)
			return nil, err
		}
	}

	return index, nil
}

func indexDirectoryExists(directoryPath string) (bool, error) {
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
