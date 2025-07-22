package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The value is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("The can go ahead by 1 step")
		fallthrough // the next case will also be executed
	case 2:
		fmt.Println("The can go ahead by 2 step")
	case 3:
		fmt.Println("The can go ahead by 3 step")
	case 4:
		fmt.Println("The can go ahead by 4 step")
	case 5:
		fmt.Println("The can go ahead by 5 step")
	case 6:
		fmt.Println("The can go ahead by 6 step")

	}
}
