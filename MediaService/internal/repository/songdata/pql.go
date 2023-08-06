package songdata

import (
	"StorageService/internal/util/apierror"
	"database/sql"
	"fmt"
)

type PqlRepository struct {
	db *sql.DB
}

func NewPqlRepository(conn *sql.DB) *PqlRepository {

	return &PqlRepository{
		db: conn,
	}
}

func (s *PqlRepository) Save(data *SongData) (*SongData, error) {

	_, err := s.db.Exec("INSERT INTO songs (file_path, title, artist) VALUES ($1, $2, $3)",
		data.FilePath, data.Title, data.Artist)
	if err != nil {
		fmt.Println("Error while saving song data:", err)
		return nil, fmt.Errorf("apierror inserting data into database: %v", err)
	}

	query := "SELECT id, file_path, title, artist FROM songs WHERE artist = $1 ORDER BY id DESC LIMIT 1"
	row := s.db.QueryRow(query, data.Artist)

	var savedData SongData
	err = row.Scan(&savedData.Id, &savedData.FilePath, &savedData.Title, &savedData.Artist)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the saved row from the database: %v", err)
	}

	return &savedData, nil
}

func (s *PqlRepository) GetById(Id int) (*SongData, error) {

	row := s.db.QueryRow("SELECT id, file_path, title, artist FROM songs WHERE id = $1", Id)

	song := &SongData{}

	err := row.Scan(&song.Id, &song.FilePath, &song.Title, &song.Artist)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apierror.NewNotFoundError("Song not found")
		}
		return nil, err
	}

	return song, nil
}
