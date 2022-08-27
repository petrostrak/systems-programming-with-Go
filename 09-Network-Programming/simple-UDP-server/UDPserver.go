package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		os.Exit(100)
	}
	PORT := ":" + arguments[1]

	s, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print("-> ", string(buffer[0:n]))

		data := []byte(buffer[0:n])
		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}
	}
}
