package app

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", getTime).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
