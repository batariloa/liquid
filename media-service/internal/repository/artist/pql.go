package artist

import (
	"StorageService/internal/util/apierror"
	"database/sql"
	"errors"
	"fmt"
)

type PqlRepository struct {
	db *sql.DB
}

func NewPqlRepository(db *sql.DB) *PqlRepository {
	return &PqlRepository{
		db: db,
	}
}

func (r *PqlRepository) GetById(artistId int) (*Artist, error) {

	query := "SELECT * FROM artists WHERE id=$1"

	row := r.db.QueryRow(query, artistId)

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

func (r *PqlRepository) Save(artist *Artist) (*Artist, error) {
	query := "INSERT INTO artists (name) VALUES($1) RETURNING id, name"

	row := r.db.QueryRow(query, artist.Name)

	var insertedArtist Artist
	err := row.Scan(&insertedArtist.ID, &insertedArtist.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to save artist: %w", err)
	}

	return &insertedArtist, nil
}
