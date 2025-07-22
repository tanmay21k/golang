package main

import "fmt"

func main() {
	fmt.Println("HELLO")
	fmt.Println("WORLD \n")

	//defer fmt.Println("BYE")
	//fmt.Println("WORLD")

	defer fmt.Println("out")
	defer fmt.Println("first")
	defer fmt.Println("in")
	defer fmt.Println("last")

}
