package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		os.Exit(1)
	}
	CONNECT := arguments[1]

	s, err := net.ResolveUDPAddr("udp", CONNECT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c, err := net.DialUDP("udp", nil, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer c.Close()

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())

	data := []byte("Hello UDP echo server!\n")
	_, err = c.Write(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("Reply: ", string(buffer[:n]))
}
