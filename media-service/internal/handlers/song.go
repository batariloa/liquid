package handlers

import (
	"StorageService/internal/apierror"
	"StorageService/internal/data"
	"StorageService/internal/handlers/helper"
	"StorageService/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleGetSongs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of songs")
}

func HandleGetSongByID(w http.ResponseWriter, r *http.Request) {

	idStr, _ := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

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

func HandleUploadSong(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) //10 MBs
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("song-file")
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

	song := data.NewSong(path, title, artistId)
	uploadedSong, err := data.SaveSong(song)

	err = service.GenerateAndPublishSongUploadEvent(uploadedSong.Id, title, artistResult.Name)
	if err != nil {
		fmt.Println("Kafka error: ", err)
		apierror.HandleAPIError(w, err)
		return
	}

	helper.WriteJSONResponse(w, uploadedSong, http.StatusCreated)
}
