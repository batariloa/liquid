package service

import (
	"StorageService/internal/repository/songdata"
	"StorageService/internal/util/apierror"
	"fmt"
)

type SongDataService struct {
	songDataRepository songdata.Repository
}

func NewSongDataService(repository songdata.Repository) *SongDataService {
	return &SongDataService{
		songDataRepository: repository,
	}
}

func (s *SongDataService) Save(data *songdata.SongData) (*songdata.SongData, error) {
	data, err := s.songDataRepository.Save(data)
	if err != nil {
		fmt.Println("Trouble saving song data", err)
		return nil, apierror.NewInternalServerError("Could not save song data.")
	}

	return data, nil
}

func (s *SongDataService) GetSongById(songId int) (*songdata.SongData, error) {

	song, err := s.songDataRepository.GetById(songId)
	if err != nil {
		return nil, err
	}

	if song == nil {
		return nil, apierror.NewNotFoundError("Song not found")
	}

	return song, nil
}
