package datastruct

type UploadKafkaEvent struct {
	ArtistName string `json:"artistName"`
	SongName   string `json:"songName"`
	SongID     int    `json:"songID"`
}
