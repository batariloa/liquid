package service

import (
	"StorageService/internal/repository/artist"
	"StorageService/internal/util/apierror"
)

type ArtistService struct {
	ArtistRepository artist.Repository
}

func NewArtistService(repository artist.Repository) *ArtistService {
	return &ArtistService{
		ArtistRepository: repository,
	}
}

func (s *ArtistService) GetArtistById(artistId int) (*artist.Artist, error) {

	result, err := s.ArtistRepository.GetById(artistId)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, apierror.NotFoundError{Message: "Artist not found."}
	}

	return result, nil
}

func (s *ArtistService) Save(artist *artist.Artist) (*artist.Artist, error) {

	res, err := s.ArtistRepository.Save(artist)
	if err != nil {
		return nil, err
	}
	return res, nil
}
