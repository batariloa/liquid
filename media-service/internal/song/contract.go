package song

type Repository interface {
	Save(data *SongData) (*SongData, error)
	GetById(Id int) (*SongData, error)
}
