package main

import "fmt"

func main() {
	var ptrDemo *int
	fmt.Println(ptrDemo)
	var number int = 52
	var ptr = &number
	fmt.Println("The address of the pointer is ", ptr)
	fmt.Println("The address of the pointer is ", *ptr)

	// Manipulating the pointer
	ptrDemo = ptr
	fmt.Println(ptrDemo) // ptrDemo -> ptr -> Address of `number`
	*ptr = *ptr * 2
	fmt.Println(*ptr)
}
