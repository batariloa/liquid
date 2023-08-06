package artist

type Repository interface {
	Save(artist *Artist) (*Artist, error)
	GetById(artistId int) (*Artist, error)
}
