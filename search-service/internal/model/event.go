package model

type SongUploadEvent struct {
	ID         int    `json:"songId"`
	Title      string `json:"title"`
	ArtistName string `json:"artistName"`
}
