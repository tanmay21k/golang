package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the rating from 0 to 5")
	// initializing the reader
	reader := bufio.NewReader(os.Stdin)
	// taking input
	inputVar, _ := reader.ReadString('\n')
	// type conversion to to float
	numRating, err := strconv.ParseFloat(strings.TrimSpace(inputVar), 64)
	if err != nil {
		fmt.Println("The error is ", err)
	} else {
		fmt.Println("The input is ", numRating+1)
		fmt.Printf("The datatype is %T \n", numRating)
	}

}
