package artist

type ArtistService interface {
	GetArtistById(artistId int) (*Artist, error)
	Save(artist *Artist) (*Artist, error)
}
