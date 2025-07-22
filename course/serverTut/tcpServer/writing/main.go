package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// listening to the specified port
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	// writint the string
	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		io.WriteString(conn, "Hello")
		fmt.Fprintln(conn, "Hows your day?")
		fmt.Fprintf(conn, "Nice talking %v", "Tanmay")
		conn.Close()
	}
}

// we need to enable the telnet services first in windows
// test the server through powershell `telnet localhost portNo`
