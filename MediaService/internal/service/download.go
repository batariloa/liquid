package service

import (
	"StorageService/internal/util/apierror"
	"os"
)

type DownloadService struct {
	songDataService *SongDataService
}

func NewDownloadService(songDataService *SongDataService) *DownloadService {

	return &DownloadService{
		songDataService: songDataService,
	}
}

func (s *DownloadService) DownloadSongById(songId int) (*os.File, error) {

	songData, err := s.songDataService.GetSongById(songId)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(songData.FilePath)
	if err != nil {
		return nil, apierror.NewNotFoundError("Song file could not be found.")
	}

	return file, nil
}
