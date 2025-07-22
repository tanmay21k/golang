package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/image.png", dogPic)
	http.ListenAndServe(":8080", nil)
}

// func dog(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html, charset ")
// 	// we are writing the file as string and not serving it from local storage
// 	io.WriteString(w, ` <img src = "https://media.istockphoto.com/id/1388420740/photo/net-zero-and-carbon-neutral-concepts-net-zero-emissions-goals-a-climate-neutral-long-term.jpg?s=612x612&w=0&k=20&c=3ZsKkJHs8FnAk5dXdCOjd85DyKu3RissYxk161yFgBM="> </img>`)
// }

// Reading from local storage
func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset = utf-8")
	io.WriteString(w, `<img src = "./image.png" </image>`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("image.png")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

// read from file and then write it using writer

// with io.copy
// with http.ServeContent & http.ServeFile
// http.FileServer **
