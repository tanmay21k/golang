package main

import "fmt"

func main() {
	var arr1 = [3]int{11, 12, 13}
	arr2 := [5]int{2: 3, 4: 10}
	var arr3 = [3]int{1, 2, 3}
	arr4 := [10]int{}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr1[2])
	fmt.Println(arr4)

}
