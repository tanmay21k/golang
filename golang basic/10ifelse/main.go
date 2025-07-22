package main

import "fmt"

func main() {
	age := 12
	//if  age := 12; age >= 18 { this can also be done specially in go!
	if age >= 18 {
		fmt.Println(age)
	} else {
		fmt.Println("minor")
	}

}
