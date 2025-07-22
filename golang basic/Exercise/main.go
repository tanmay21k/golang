// 1. Initialize a SLICE of int using a composite literal; print out the slice; range over the slice
// printing out just the index; range over the slice printing out both the index and the value

// package main

// import "fmt"

// func main() {
// 	sl1 := []int{1, 2, 3}
// 	// fmt.Println(sl1)
// 	for _, val := range sl1 {
// 		fmt.Printf("The index is index and the value is %v \n", val)
// 		// fmt.Println("The value is ", val)
// 	}
// }

// -----------------------

// 2. Initialize a MAP using a composite literal where the key is a string and the value is an int;
// print out the map; range over the map printing out just the key; range over the map
// printing out both the key and the value

// package main

// import "fmt"

// func main() {
// 	map1 := make(map[string]int)
// 	map1["index1"] = 1
// 	map1["index2"] = 2
// 	map1["index3"] = 3
// 	map1["index4"] = 4
// 	map1["index5"] = 5
// 	map1["index6"] = 6

// 	for index, val := range map1 {
// 		fmt.Printf("Current index is %v and the value is %v \n", index, val)
// 	}

// }

// -------------------

// Create a new type called person which is STRUCT that stores fName and lName

// package main

// import "fmt"

// type person struct {
// 	fName string
// 	lName string
// }

// func main() {
// 	person1 := person{"Firstname", "Lastname"}
// 	fmt.Printf("%#v", person1)
// }

// Take the STRUCT “person” in the previous exercise and add a field called “favFood” that
// stores a slice of string; for the variable “p1” use a composite literal to add values to the
// field favFood; print out the values in favFood; also print out the values for “favFood” by
// using a for range loop

// package main

// import "fmt"

// type person struct {
// 	favFood []string
// 	pName   string
// }

// func main() {
// 	p1 := person{
// 		[]string{"apple", "banana"},
// 		"name",
// 	}

//		// favFood := []string{"apple", "banana"}
//		fmt.Printf("%#v \n", p1)
//		fmt.Println(p1.pName)
//		// fmt.Println(p1.favFood)
//		for _, val := range p1.favFood {
//			fmt.Println(val)
//		}
//	}

// Create a method walk which is executed via the struct

// package main

// import "fmt"

// type person struct {
// 	pName string
// }

// func (p person) walk() string {
// 	return p.pName
// }

// func main() {
// 	p1 := person{
// 		"PERSON1",
// 	}
// 	obj := p1.walk()
// 	fmt.Println(obj)
// }

// ------------------------

// 8. Create a struct called sedan and truck and it should contain
// sedan -> luxury
// truck -> fourWheel
// and vehicle struct

// Create struct named vehicle
// vehicle should contain doors color

// package main

// import "fmt"

// type vehicle struct {
// 	doors int
// 	color string
// }

// type sedan struct {
// 	luxury bool
// 	vehicle
// }

// type truck struct {
// 	fourWheel bool
// 	vehicle
// }

// func main() {
// 	sedan1 := truck{
// 		true,
// 		vehicle{
// 			4,
// 			"red",
// 		},
// 	}
// 	fmt.Println(sedan1)

// 	truck1 := truck{
// 		false,
// 		vehicle{
// 			2,
// 			"Green",
// 		},
// 	}
// 	fmt.Println(truck1)

// 	v1 := vehicle{
// 		4,
// 		"black",
// 	}
// 	sedan2 := sedan{
// 		true,
// 		v1,
// 	}
// 	fmt.Println(sedan2.color)
// 	fmt.Println(sedan2.luxury)

// func (t sedan) transportationDevice() string {
// 	return fmt.Println("The color of the vehicle is ",sedan2.color)
// }
// fmt.Println(sedan2.transportationDevice())
// }

// Create a custom datatype of int type
// check whether implicit type conversion is possible

package main

import "fmt"

type customType int

func main() {
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)
	fmt.Println(string(s[:14]))
	fmt.Println(string(s[10:22]))
	sByte := []byte(s)
	fmt.Printf("%T\n ", sByte)
	fmt.Println(sByte) // string conversion to byte
	fmt.Println(sByte[0])
	sNew := string(sByte)
	fmt.Println(sNew)

	var num1 customType = 21
	defer fmt.Printf("%T\n ", num1)
	num1new := int(num1)
	defer fmt.Printf("%T\n ", num1new)
}
