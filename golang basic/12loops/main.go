package main

import "fmt"

func main() {

	days := map[int]string{
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
		4: "Thursday",
		5: "Friday",
	}
	for d := 0; d < len(days); d++ { // we are going from index 0 to 4 here, 0 is not defined as key and 5 : friday will not be reached
		fmt.Println(days[d])
	}

	fmt.Printf("\n Loop using the range \n")
	for i := range days { // similar to -> for i in range () of python
		fmt.Println(days[i])
	}

	// IMP -> i resembles the index here.

	for no, day := range days { // here no is the index
		fmt.Printf("The day number is %v and the day is %v \n", no, day)
	}

	for _, day := range days {
		fmt.Printf("The day number is v and the day is %v \n", day)
	}
}
