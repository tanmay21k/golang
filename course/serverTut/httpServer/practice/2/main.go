package main

import (
	"html/template"
	"net/http"
)

func serveTemp(W http.ResponseWriter, R *http.Request) {
	// Parse the template file
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}

	// data := "Tanmay"
	// Execute the template with the data
	err = tpl.ExecuteTemplate(W, "index.gohtml", "Tanmay")
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", serveTemp)
	http.ListenAndServe(":8080", nil)
}
