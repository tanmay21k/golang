// Package constantpkg demonstrates Go constant declaration patterns.
// It shows how constants work with type declarations and implicit value repetition.
package constantpkg

import "fmt"

// README contains good article

// When a const initializer is omitted, Go repeats the previous expression.
// The blank identifier (_) also receives that expression, but its value is discarded.

type weekday int
const ( 
Sunday weekday = 1
_
Monday 
Tuesday weekday = 122
Wednesday
)


// ConstantPrinter demonstrates how constants behave with type declarations
// and implicit value repetition using the blank identifier.
func ConstantPrinter() {
	fmt.Println(Sunday) // 1
	fmt.Println(Monday) // 1
	fmt.Println(Tuesday) // 122
	fmt.Println(Wednesday) // 122
}