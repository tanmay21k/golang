package main

import (
	"fmt"
	"net/url"
)

const MyUrl string = "https://pkg.go.dev/github.com/cpmech/gosl/io%23ReadFile?param1=value1&param2=value2&param3=value3"

func main() {
	fmt.Println("Dealing with parameters of URL")

	// parsing the url
	result, err := url.Parse(MyUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", result.Scheme)     // https
	fmt.Println("Host:", result.Host)         // pkg.go.dev
	fmt.Println("Path:", result.Path)         // /github.com/cpmech/gosl/io#ReadFile
	fmt.Println("Port:", result.Port())       // (empty, since no port is specified)
	fmt.Println("RawQuery:", result.RawQuery) // param1=value1&param2=value2
	fmt.Println("RawPath:", result.RawPath)   // /github.com/cpmech/gosl/io%23ReadFile

	qparams := result.Query()
	fmt.Printf("The type of the parameters is %T \n", qparams)
	fmt.Println("The value of key param1 is", qparams["param1"])
	fmt.Println(qparams)

	// iterating over the params
	for index, val := range qparams {
		fmt.Println(index, "stores", val)
	}
}
