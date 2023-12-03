package types

type SongFile struct {
	Path     string `json:"path"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
}
