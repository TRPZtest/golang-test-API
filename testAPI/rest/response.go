package respone

import (
	"encoding/json"
	"log"
	"net/http"
)

func InternalServerErrorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func BadRequestResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func SuccessJsonResponse(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
