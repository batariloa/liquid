package model

type Song struct {
	ID         string `json:"songId"`
	Title      string `json:"title"`
	ArtistName string `json:"artistName"`
}
