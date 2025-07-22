package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	router1 := mux.NewRouter()                         // creating router
	router1.HandleFunc("/", serverHome).Methods("GET") // handling functions at different routes
	log.Fatal(http.ListenAndServe(":4000", router1))   // requires port and router
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to golang docs by Tanmay</h1>"))
}
