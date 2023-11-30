package data

import (
	"StorageService/internal/db"
	"StorageService/internal/util/apierror"
	"database/sql"
	"fmt"
)

type SongData struct {
	Id       int    `json:"id"`
	FilePath string `json:"file_path"`
	Title    string `json:"title"`
	Artist   int    `json:"artist"`
}

func SaveSong(data *SongData) (*SongData, error) {

	_, err := db.DB.Exec("INSERT INTO songs (file_path, title, artist) VALUES ($1, $2, $3)",
		data.FilePath, data.Title, data.Artist)
	if err != nil {
		fmt.Println("Error while saving song data:", err)
		return nil, fmt.Errorf("apierror inserting data into database: %v", err)
	}

	query := "SELECT id, file_path, title, artist FROM songs WHERE artist = $1 ORDER BY id DESC LIMIT 1"
	row := db.DB.QueryRow(query, data.Artist)

	var savedData SongData
	err = row.Scan(&savedData.Id, &savedData.FilePath, &savedData.Title, &savedData.Artist)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the saved row from the database: %v", err)
	}

	return &savedData, nil
}

func GetSongById(Id int) (*SongData, error) {

	row := db.DB.QueryRow("SELECT id, file_path, title, artist FROM songs WHERE id = $1", Id)

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
