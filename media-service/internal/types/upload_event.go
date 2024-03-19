package types

type UploadSongEvent struct {
	ArtistName string `json:"artistName"`
	Title      string `json:"title"`
	SongID     int    `json:"songId"`
	UserID     int    `json:"userId"`
}
