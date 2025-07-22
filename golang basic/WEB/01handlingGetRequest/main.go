package main

import (
	"fmt"
	"io"
	"net/http"
)

var url string = "https://www.wikipedia.org"

func main() {
	fmt.Println("Web request")
	//getting the response
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // always remember to close

	fmt.Printf("The response is of type %T \n", response)
	// dont go -> fmt.Println(response.Body)

	// reading the response
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
