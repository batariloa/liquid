package handlers

import (
	"StorageService/internal/apierror"
	"StorageService/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (*Handler) HandleDownloadSong(w http.ResponseWriter, r *http.Request) {

	idStr, _ := mux.Vars(r)["id"] // NOTE: Safe to ignore error, because it's always defined.

	id, err := strconv.Atoi(idStr)
	if err != nil {
		apierror.HandleAPIError(w, apierror.NewBadRequestError("Please provide a valid ID"))
		return
	}

	if err != nil {
		fmt.Println("Err in download handler", err)
		apierror.HandleAPIError(w, err)
		return
	}

	file, err := service.DownloadSongById(id)
	if err != nil {
		fmt.Println("Err in download handler", err)
		apierror.HandleAPIError(w, err)
		return
	}

	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeContent(w, r, file.Name(), time.Time{}, file)
}
