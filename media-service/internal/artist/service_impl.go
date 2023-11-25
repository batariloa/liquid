package artist

import (
	"StorageService/internal/util/apierror"
)

type ArtistServiceImpl struct {
	ArtistRepository Repository
}

func NewArtistService(repository Repository) *ArtistServiceImpl {
	return &ArtistServiceImpl{
		ArtistRepository: repository,
	}
}

func (s *ArtistServiceImpl) GetArtistById(artistId int) (*Artist, error) {

	result, err := s.ArtistRepository.GetById(artistId)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, apierror.NotFoundError{Message: "Artist not found."}
	}

	return result, nil
}

func (s *ArtistServiceImpl) Save(artist *Artist) (*Artist, error) {

	res, err := s.ArtistRepository.Save(artist)
	if err != nil {
		return nil, err
	}
	return res, nil
}
