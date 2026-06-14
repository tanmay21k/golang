// Package iotapkg demonstrates the use of iota for creating enumerations in Go.
// It shows various patterns of using iota for constant value assignment.
package iotapkg

import "fmt"

type weekday int

// The code below demonstrates usagage of iota and const to perform auto-incrementing
// In case of any doubt, refer the constantpkg

const (
	Sunday weekday = iota + 1 // iota(0) + 1 --> 1
	// commented line
	_ // iota = 2
	Monday // iota = 3
	Tuesday weekday = 231 // iota = 4
	_ // iota = 5
	Wednesday // iota = 6
	Thursday weekday = iota // iota  = 7
)

// IotaConstant demonstrates iota usage within a const block.
func IotaConstant() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
}

// IotaOutsideConst demonstrates that iota is reset to 0 outside of a const block.
func IotaOutsideConst() {
	// since this is not in const block iota will be 0 and will not increment
	const num1 int = iota * 1;
	num2 := 123
	num3 := 2312
	const num4 = iota

	fmt.Printf("Initially num1 is initiated with iota has value %v \n", num1)
	fmt.Println("Filler numbers:", num2 , " & ",num3 )
	fmt.Printf("Now the value for num4 to %v", num4) 
}