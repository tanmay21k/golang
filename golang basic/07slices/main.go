package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"apple", "mango", "jackfruit"}
	fruitList = append(fruitList[1:])
	fmt.Println("The slice is ", fruitList)

	fruitList = append(fruitList, "Fruit1", "Fruit2")
	fmt.Println("The slice is ", fruitList)

	fruitList = append(fruitList[1:2])
	fmt.Println("The slice is ", fruitList)

	var highScores = make([]int, 4)
	highScores[0] = 32
	highScores[1] = 31
	highScores[2] = 22
	highScores[3] = 11

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

}
