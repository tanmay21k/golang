// had to create the dependencies via the go mod inti filename
// then had to create a build fileName.exe
// visit the port on local host to get the data on terminal
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Adding a timeout
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		panic(err)
	}
	// using NewScanner
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard what you said %v", ln) // writing in the terminal
	}
	defer conn.Close()

	fmt.Println("***END CODE***")
}
