package handlers

import (
	"StorageService/internal/apierror"
	"StorageService/internal/auth"
	"StorageService/internal/data"
	"StorageService/internal/service"
	"StorageService/internal/types"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (*Handler) HandleGetSongByID(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	id, err := strconv.Atoi(idStr)
	if err != nil {
		apierror.HandleAPIError(w, apierror.NewBadRequestError("Please provide a valid ID"))
		return
	}

	song, err := data.GetSongById(id)
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

func (h *Handler) HandleUploadSong(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.ExtractUserIdFromToken(r)
	if err != nil {
		apierror.HandleAPIError(w, apierror.NewUnauthorizedRequestError("User ID not found in request."))
		return
	}

	log.Println("Got User ID from Token:", userId)

	err = r.ParseMultipartForm(10 << 20) //10 MBs
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("song-file")
	if err != nil {
		log.Println("apierror retrieving file", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	artistId, err := getArtistID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title, err := getSongTitle(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	artistResult, err := data.GetArtistById(artistId)
	if err != nil {
		apierror.HandleAPIError(w, err)
		return
	}

	path, err := service.UploadSong(&file, handler)
	if err != nil {
		fmt.Println("error uploading song")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	song := data.NewSong(path, title, artistId, userId)
	uploadedSong, err := data.SaveSong(song)

	event := types.UploadSongEvent{
		SongID:     uploadedSong.Id,
		Title:      title,
		ArtistName: artistResult.Name,
		UserID:     userId,
	}

	err = service.GenerateAndPublishSongUploadEvent(event, h.EventPublisher)
	if err != nil {
		fmt.Println("Kafka error: ", err)
		apierror.HandleAPIError(w, err)
		return
	}

	WriteJSONResponse(w, uploadedSong, http.StatusCreated)
}

func getArtistID(r *http.Request) (int, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MBs
	if err != nil {
		return 0, apierror.NewBadRequestError("Error parsing multipart form")
	}

	artistID, err := strconv.Atoi(r.FormValue("artistId"))
	if err != nil {
		return 0, apierror.NewBadRequestError("Artist ID not provided or invalid")
	}

	return artistID, nil
}

func getSongTitle(r *http.Request) (string, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MBs
	if err != nil {
		return "", apierror.NewBadRequestError("Error parsing multipart form")
	}

	title := r.FormValue("title")
	if title == "" {
		return "", apierror.NewBadRequestError("Song title not provided or invalid")
	}

	return title, nil
}

func getSongId(r *http.Request) (string, error) {
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return "", apierror.NewBadRequestError("Error decoding request body")
	}

	id, ok := requestBody["songId"].(string)
	if !ok {
		return "", apierror.NewBadRequestError("Song ID not provided or invalid")
	}

	return id, nil
}
