package handlers

import (
	"StorageService/internal/artist"
	"StorageService/internal/handlers/helper"
	"StorageService/internal/kafka"
	"StorageService/internal/song"
	"StorageService/internal/upload"
	"StorageService/internal/util/apierror"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SongHandler struct {
	UploadService   upload.UploadService
	SongDataService song.SongDataService
	KafkaService    kafka.KafkaService
	ArtistService   artist.ArtistService
}

func NewSongHandler(
	uploadService upload.UploadService,
	songDataService song.SongDataService,
	kafkaService kafka.KafkaService,
	artistService artist.ArtistService,
) *SongHandler {

	return &SongHandler{
		UploadService:   uploadService,
		SongDataService: songDataService,
		KafkaService:    kafkaService,
		ArtistService:   artistService,
	}
}

func (h *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "List of songs")
}

func (h *SongHandler) GetSongByID(w http.ResponseWriter, r *http.Request) {

	idStr, _ := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	id, err := strconv.Atoi(idStr)
	if err != nil {
		apierror.HandleAPIError(w, apierror.NewBadRequestError("Please provide a valid ID"))
		return
	}

	song, err := h.SongDataService.GetSongById(id)
	if err != nil {
		apierror.HandleAPIError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(song)
	if err != nil {
		apierror.HandleAPIError(w, err)
		return
	}
}

func (h *SongHandler) UploadSong(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) //10 MBs

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("song_file")
	if err != nil {
		log.Println("apierror retrieving file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	artistId, err := helper.GetArtistID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title, err := helper.GetSongTitle(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	artistResult, err := h.ArtistService.GetArtistById(artistId)
	if err != nil {
		apierror.HandleAPIError(w, err)
		return
	}

	path, err := h.UploadService.UploadSong(&file, handler)
	if err != nil {
		fmt.Println("error uploading song")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	song := song.New(path, artistId, title)
	uploadedSong, err := h.SongDataService.Save(song)

	err = h.UploadService.GenerateAndPublishSongUploadEvent(uploadedSong.Id, title, artistResult.Name)
	if err != nil {
		fmt.Println("Kafka error: ", err)
		apierror.HandleAPIError(w, err)
		return
	}

	helper.WriteJSONResponse(w, uploadedSong, http.StatusCreated)
}
