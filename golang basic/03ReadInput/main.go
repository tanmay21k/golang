// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Enter the rating")
// 	reader := bufio.NewReader(os.Stdin)
// 	inputVar, _ := reader.ReadString('\n')
// 	fmt.Println("The rating is", inputVar)
// 	fmt.Printf("The type of the data is %T", inputVar)
// }

package main

import (
	"fmt"
	"strconv"
)

func main() {

	string1 := "32.22"
	inNumber, err := strconv.ParseFloat(string1, 64)
	if err != nil {
		fmt.Println("Error was", err)
		return
	}
	fmt.Printf("The type after parsing to float is %T \n", inNumber)
	fmt.Println("The value is", inNumber)

}
