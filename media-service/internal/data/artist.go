package data

import (
	"StorageService/internal/apierror"
	"StorageService/internal/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetArtistById(artistId int) (*Artist, error) {

	query := "SELECT * FROM artists WHERE id=$1"

	row := db.DB.QueryRow(query, artistId)

	var resultArtist Artist
	err := row.Scan(&resultArtist.ID, &resultArtist.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apierror.NewNotFoundError("Artist not found.")
		}
		return nil, fmt.Errorf("failed to fetch artist: %w", err)
	}

	return &resultArtist, nil
}

func SaveArtist(artist *Artist) (*Artist, error) {
	query := "INSERT INTO artists (name) VALUES($1) RETURNING id, name"

	row := db.DB.QueryRow(query, artist.Name)

	var insertedArtist Artist
	err := row.Scan(&insertedArtist.ID, &insertedArtist.Name)
	if err != nil {
		log.Print("Failed to save artist", err)
		return nil, fmt.Errorf("failed to save artist: %w", err)
	}

	return &insertedArtist, nil
}
