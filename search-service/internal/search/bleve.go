package search

import (
	"log"
	"os"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/lang/en"
)

const indexDirectory = "songs.bleve"

// initializes or opens the Bleve index.
func InitializeBleveIndex() (bleve.Index, error) {

	indexExists, err := indexDirectoryExists(indexDirectory)
	if err != nil {
		log.Printf("Failed to check if index directory exists: %v", err)
		return nil, err
	}

	if indexExists {
		log.Println("Opening existing Bleve index.")
		return bleve.Open(indexDirectory)
	}

	log.Println("Creating a new Bleve index.")
	return createNewIndex()
}

func createNewIndex() (bleve.Index, error) {
	englishTextFieldMapping := bleve.NewTextFieldMapping()
	englishTextFieldMapping.Analyzer = en.AnalyzerName

	songMapping := bleve.NewDocumentMapping()
	songMapping.AddFieldMappingsAt("title", englishTextFieldMapping)
	songMapping.AddFieldMappingsAt("artist", englishTextFieldMapping)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.TypeField = "type"
	indexMapping.DefaultAnalyzer = "en"
	indexMapping.AddDocumentMapping("song", songMapping)

	return bleve.New(indexDirectory, indexMapping)
}

func indexDirectoryExists(directoryPath string) (bool, error) {
	if _, err := os.Stat(directoryPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
