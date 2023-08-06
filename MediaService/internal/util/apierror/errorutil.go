package apierror

import (
	"net/http"
)

func HandleAPIError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case NotFoundError:
		handleNotFoundError(w, err)
	case InternalServerError:
		handleInternalServerError(w, err)
	case BadRequestError:
		handleBadRequestError(w, err)
	default:
		handleDefaultError(w, err)
	}
}

func handleNotFoundError(w http.ResponseWriter, err error) {
	apiErr := err.(NotFoundError)
	http.Error(w, apiErr.Message, http.StatusNotFound)
}

func handleInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, "Something went wrong.", http.StatusInternalServerError)
}

func handleBadRequestError(w http.ResponseWriter, err error) {
	apiErr := err.(BadRequestError)
	http.Error(w, apiErr.Message, http.StatusBadRequest)
}

func handleDefaultError(w http.ResponseWriter, err error) {
	http.Error(w, "There was an error processing your request.", http.StatusInternalServerError)
}
