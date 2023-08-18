package datastruct

type UploadKafkaEvent struct {
	ArtistName string `json:"artistName"`
	SongTitle  string `json:"title"`
	SongID     int    `json:"songId"`
}
