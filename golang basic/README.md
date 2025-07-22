# Go Programming Language – Comprehensive Guide  

Go (or Golang) is a statically typed, compiled programming language designed at Google. It is known for its simplicity, efficiency, and excellent support for concurrency. Unlike Java or C++, Go does not use a Virtual Machine (VM), making it lightweight and efficient.  

---

# **Running a Go Program**  
### **Execute a Go program using**:  
```sh
go run filename.go
```  
This compiles and executes the Go file but does not create an executable file.  

### **Compile a Go program into a binary executable**:  
```sh
go build filename.go
```  
This generates a binary file (`.exe` on Windows, or a standard binary file on Linux/macOS), which can be executed directly.  

---

# **Go Modules & Dependency Management**  
Go modules help manage dependencies similar to `npm init` in JavaScript.  
### **Initialize a new Go module**:  
```sh
go mod init module_name
```  
This command creates a `go.mod` file, which keeps track of dependencies.  

### **Fetching dependencies**:  
Use `go get` to fetch external packages:  
```sh
go get -u github.com/gorilla/mux
```
  
### **Useful Go module commands**:  
- `go mod tidy` → Cleans unused dependencies.  
- `go mod verify` → Verifies dependencies for integrity.  
- `go list -m all` → Lists all module dependencies.  
- `go mod vendor` → Creates a local `vendor/` directory with dependencies.  

---

# **Go Help System**  
To get help on Go commands and features:  
```sh
go help
```
You can also visit [https://pkg.go.dev](https://pkg.go.dev) to read about Go packages.  

---

# **Go Program Structure & Entry Point**  
In Go, the entry point of a program is the `main` function in the `main` package.  

Example:  
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

---

# **`fmt` Package – Printing & Formatting**  
The `fmt` package provides functions for printing formatted text.  

- **Basic printing**:  
  ```go
  fmt.Println("Hello, world!")
  ```
- **Formatted printing with placeholders (`%T`, `%v`, etc.)**:  
  ```go
  name := "Alice"
  age := 25
  fmt.Printf("Name: %s, Age: %d\n", name, age)
  ```
- **Placeholders in `fmt.Printf`**:  
  - `%d` → Integer  
  - `%s` → String  
  - `%f` → Float  
  - `%t` → Boolean  
  - `%T` → Type of the variable  
  - `%+v` → Print struct with field names  

---

# **Data Types in Go**  

Go has primitive and complex data types:  

### **Primitive types**  
- `bool` → Boolean values (`true` or `false`)  
- `string` → Text data  
- `int`, `int8`, `int16`, `int32`, `int64` → Integer types  
- `uint`, `uint8`, `uint16`, `uint32`, `uint64` → Unsigned integers  
- `float32`, `float64` → Floating-point numbers  
- `complex64`, `complex128` → Complex numbers  

### **Complex types**  
- `array` → Fixed-size list of elements  
- `slice` → Dynamic-size list  
- `map` → Key-value pairs (like a dictionary)  
- `struct` → Custom data structure  
- `pointer` → Stores memory addresses  
- `channel` → Used in concurrency  

---

# **Variables & Constants**  

### **Variable Declaration**  
```go
var name string = "Alice"
var age int = 25
```
Go also supports type inference:  
```go
name := "Alice"  // Automatically inferred as string
age := 25        // Automatically inferred as int
```

### **Constants**  
Constants are declared using `const`:  
```go
const pi float64 = 3.14159
```

---

# **Taking User Input**  

Use `bufio.NewReader` and `os.Stdin` for input handling:  
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello,", input)
}
```
**Note:** `_` is used to ignore errors if they are not needed.  

---

# **`strconv` Package – String Conversions**  

Convert strings to numbers using `strconv`:  
```go
import "strconv"

str := "42"
num, _ := strconv.Atoi(str) // Converts string to integer
fmt.Println(num)
```

Convert numbers to strings:  
```go
num := 42
str := strconv.Itoa(num) // Converts integer to string
fmt.Println(str)
```

---

# **Go Arrays & Slices**  

### **Array** (Fixed size)  
```go
var arr = [3]int{1, 2, 3}
fmt.Println(arr)
```

### **Slice** (Dynamic size)  
```go
var slice = []int{1, 2, 3}
slice = append(slice, 4) // Append element
fmt.Println(slice)
```

### **Slicing an array**  
```go
newSlice := arr[1:3] // Includes index 1, excludes index 3
```

### **Sorting a Slice**  
```go
import "sort"

numbers := []int{3, 1, 4, 1, 5}
sort.Ints(numbers)
fmt.Println(numbers)
```

---

# **Go Maps (HashMap equivalent)**  

Declaration:  
```go
var myMap = map[string]int{"Alice": 25, "Bob": 30}
```

Adding & Deleting:  
```go
myMap["Charlie"] = 35
delete(myMap, "Bob")
```

Iterating over a map:  
```go
for key, value := range myMap {
    fmt.Println(key, value)
}
```

---

# **Structs – Custom Data Types**  

```go
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	user := User{"Alice", 30, "alice@example.com"}
	fmt.Println(user)
}
```

Printing a struct in detail:  
```go
fmt.Printf("%+v\n", user)
```

---

# **Functions & Methods in Go**  

### **Basic Function**  
```go
func add(a int, b int) int {
    return a + b
}
```

### **Function with Multiple Return Values**  
```go
func swap(x, y string) (string, string) {
    return y, x
}
```

### **Methods (Functions tied to Structs)**  
```go
type User struct {
    Name string
}

func (u User) greet() {
    fmt.Println("Hello,", u.Name)
}
```

---

# **Concurrency in Go – Goroutines**  
Goroutines enable concurrent execution.  
```go
func printMessage(msg string) {
    fmt.Println(msg)
}

func main() {
    go printMessage("Hello")
    go printMessage("World")
}
```
This executes `printMessage` in separate lightweight threads.

---

# **Go Web Development – HTTP Requests**  

Using `net/http` for GET requests:  
```go
import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```

---

# **Go Routing – Gorilla Mux**  
Install using:  
```sh
go get -u github.com/gorilla/mux
```

Basic routing example:  
```go
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Go Web!")
	}).Methods("GET")

	http.ListenAndServe(":8080", router)
}
```

---


Yes! You covered a lot, but there are still a few more Go basics and concepts that are useful to know. Here are some additional topics you might want to include:

---

## **Go Variable Declarations and Shadowing**
In Go, variables can be declared in multiple ways:

```go
var name string = "John"
var age int = 30
isAlive := true // Short declaration (only inside functions)
```

### **Variable Shadowing**
- In Go, variables declared in an inner scope can **shadow** variables from an outer scope.

```go
package main

import "fmt"

func main() {
	name := "John"
	fmt.Println(name) // John

	if true {
		name := "Doe" // Shadows outer 'name'
		fmt.Println(name) // Doe
	}

	fmt.Println(name) // John (Outer scope remains unchanged)
}
```

---

## **Loops in Go (Expanded)**

Go only has `for` loops, but they can be used in different ways:

1. **Basic Loop**
```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

2. **While-like Loop**
```go
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}
```

3. **Infinite Loop**
```go
for {
	fmt.Println("This is an infinite loop")
}
```

4. **Looping Through Arrays/Slices/Maps**
```go
arr := []string{"Go", "Python", "Java"}

for index, value := range arr {
	fmt.Println(index, value)
}

myMap := map[string]int{"one": 1, "two": 2}
for key, value := range myMap {
	fmt.Println(key, value)
}
```

---

## **Errors in Go**
Go does not have exceptions like other languages (no `try-catch`). Instead, it handles errors using multiple return values.

### **Error Handling Example**
```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
```

### **Using `fmt.Errorf` for Formatting Errors**
```go
import "fmt"

func main() {
	err := fmt.Errorf("an error occurred: %v", "something went wrong")
	fmt.Println(err)
}
```

---

## **Concurrency in Go (Goroutines & Channels)**

### **Goroutines**
A goroutine is a lightweight thread managed by Go.

```go
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go printMessage("Hello from Goroutine") // Runs in the background
	printMessage("Hello from Main")         // Runs in the main thread
}
```

### **Channels**
Goroutines communicate using **channels**.

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!"
	}()

	msg := <-ch // Receiving value from channel
	fmt.Println(msg)
}
```

---

## **Select Statement (Handling Multiple Channels)**
`select` allows a goroutine to wait on multiple communication operations.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout reached")
	}
}
```

---

## **Go Panic and Recover**
### **Panic**
`panic` is used to stop the execution when something goes terribly wrong.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("This will not execute if panic is not recovered")

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic") // This will not execute
}
```

### **Recover from Panic**
`recover` is used to regain control from a panic.

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic") // Will not execute
}
```

---

## **Go Context Package (Managing Goroutines)**
The `context` package is used to control goroutines efficiently.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Process stopped!")
			return
		default:
			fmt.Println("Processing...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go process(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("Main function completed")
}
```

---

## **Reflection in Go**
Reflection allows you to inspect variables, types, and structures at runtime.

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myVar int = 42
	fmt.Println("Type of myVar:", reflect.TypeOf(myVar))
}
```

---

## **More on Go Maps**
- **Checking if a Key Exists**
```go
value, exists := myMap["key"]
if exists {
	fmt.Println("Key exists:", value)
} else {
	fmt.Println("Key does not exist")
}
```

---

## **More String Functions (strings package)**
The `strings` package provides useful functions:

```go
import "strings"

str := "Hello, World"
fmt.Println(strings.ToLower(str))  // "hello, world"
fmt.Println(strings.ToUpper(str))  // "HELLO, WORLD"
fmt.Println(strings.HasPrefix(str, "Hello")) // true
fmt.Println(strings.HasSuffix(str, "World")) // true
fmt.Println(strings.Contains(str, "lo"))     // true
fmt.Println(strings.Split(str, ", "))        // ["Hello" "World"]
```

---

## **Environment Variables in Go**
You can use `os` package to read environment variables.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("MY_VAR", "Golang")
	fmt.Println("Environment variable:", os.Getenv("MY_VAR"))
}
```

---

## **Logging in Go**
The `log` package is used for logging messages.

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, _ := os.Create("app.log")
	log.SetOutput(file)

	log.Println("This is a log message")
	log.Fatal("This is a fatal error") // Terminates the program
}
```

---

## **Go Worker Pools**
Worker pools allow you to control the number of goroutines executing tasks.

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
}
```

---

### That’s all for now! 😃 Hope this helps! 🚀 Let me know if you want more details on any topic.