package main

import "fmt"

const isExample bool = true
const IsGlobal = true // Capital letter is used to assign Global scoped variables.

func main() {
	var username string = "Tanmay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var number int = 23
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	var smallNo uint8 = 255
	fmt.Println(smallNo)
	fmt.Printf("Variable is of type: %T \n", smallNo)

	var isLogged = true // Lexer does the assigning if the datatype is not mentioned
	fmt.Printf("Variable is of type: %T \n", isLogged)

	age := 2
	age = 32
	fmt.Println(age)

	fmt.Println(isExample)
	fmt.Printf("Variable is of type: %T \n", isExample)
}
