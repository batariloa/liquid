package mapper

import (
	"strconv"

	"github.com/batariloa/search-service/internal/model"
)

func EventToSong(e *model.SongUploadEvent) *model.Song {

	return &model.Song{

		ID:         strconv.Itoa(e.ID),
		Title:      e.Title,
		ArtistName: e.ArtistName,
	}
}
