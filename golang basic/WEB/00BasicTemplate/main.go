package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // serving the path
		json.NewEncoder(w).Encode("Hello World") // for creating json data
	}).Methods("Get")

	http.ListenAndServe(":4001", router) // creating port
}
