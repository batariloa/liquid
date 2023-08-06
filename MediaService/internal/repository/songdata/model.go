package songdata

type SongData struct {
	Id       int    `json:"id"`
	FilePath string `json:"file_path"`
	Title    string `json:"title"`
	Artist   int    `json:"artist"`
}

func New(FilePath string, ArtistId int, Title string) *SongData {

	return &SongData{
		FilePath: FilePath,
		Title:    Title,
		Artist:   ArtistId,
	}
}
