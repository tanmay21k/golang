// Implementing using a interface

package main

import (
	"fmt"
	"net/http"
)

type hotdog int // creating interface

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code from the func")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8000", d) // using inteface to execute the listen and serve since any interface is doable
}
