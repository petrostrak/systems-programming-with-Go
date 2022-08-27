package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(100)
	}
	SERVER := "localhost" + ":" + arguments[1]

	s, err := net.ResolveTCPAddr("tcp", SERVER)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)

	for {
		// Get and process client requests.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print("> ", string(buffer[0:n]))

		_, err = conn.Write(buffer)

		conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
