package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// Reads the data from a socket file using Read()
func readSocket(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("-> ", string(buf[0:n]))
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a socket file.")
		os.Exit(100)
	}
	socketFile := arguments[1]

	// The net.Dial() function with the right first argument allows you
	// to connect to the socket file before you try to read from it.
	c, err := net.Dial("unix", socketFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	go readSocket(c)
	for {
		_, err := c.Write([]byte("Hello server!"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		time.Sleep(time.Second)
	}
}
