package main

import (
	"fmt"
	"net/http"
	"strings"
)

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello world")
}

func Dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Dog goes bhaw bhaw")
}

func NameFunc(res http.ResponseWriter, req *http.Request) {
	result := strings.Split(req.URL.Path, "/")
	if len(result) > 2 {
		name := result[2]
		fmt.Fprintf(res, "Hello %s", name)
	} else {
		fmt.Fprintf(res, "No name provided!")
	}
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/dog", Dog)
	http.HandleFunc("/me/", NameFunc)
	http.ListenAndServe(":8080", nil)
}
