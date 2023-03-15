package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	handlers "testAPI/handlers"
	"testAPI/models"

	"github.com/gorilla/mux"
)

func POST(w http.ResponseWriter, r *http.Request) {

	var requestBody models.GetMaxPriceRequest

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &requestBody)

	w.WriteHeader(http.StatusNoContent)
}

func getMaxPrice(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	handlerPOST := handlers.NewGetMaxPriceHandlerPOST()
	handlerGET := handlers.NewGetMaxPriceHandlerGET()

	router.HandleFunc("/", handlerPOST.Execute).Methods("POST")
	router.HandleFunc("/", handlerGET.Execute).Methods("GET")
	// router.HandleFunc("/", homeLink).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
