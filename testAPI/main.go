package main

import (
	"log"
	"net/http"
	handlers "testAPI/handlers"

	"github.com/gorilla/mux"
)

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
