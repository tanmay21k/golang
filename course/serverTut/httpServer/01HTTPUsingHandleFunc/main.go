// // Implementing using function handler
// package main

// import (
// 	"io"
// 	"net/http"
// )

// func d(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "dog dog dog")
// }

// func c(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "cat cat cat")
// }

//	func main() {
//		http.HandleFunc("/cat", c)
//		http.HandleFunc("/dog", d)
//		http.ListenAndServe(":8000", nil)
//	}
package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Home is the handler for the "/" route
func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "HOME PAGE")
}

// dog is the handler for the "/dog" route
func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "bha bhau bhau")
}

// me is the handler for the "/me" route, which should print the name
func me(res http.ResponseWriter, req *http.Request) {
	// Extract the 'name' from the URL path ("/me/{name}")
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) > 2 {
		name := parts[2]
		fmt.Fprintf(res, "Hello, %s!", name)
	} else {
		fmt.Fprintf(res, "No name provided!")
	}
}

func main() {
	// Register the route handlers
	http.HandleFunc("/", Home)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
