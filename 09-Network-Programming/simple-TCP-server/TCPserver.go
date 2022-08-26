// A TCP server that implements the Echo service. The Echo service
// is usually implemented using the UDP protocol due to its simplicity,
// but it can also be implemented with TCP.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		os.Exit(100)
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()

	// Only after a successful call to Accept() , the TCP server
	// can start interacting with TCP clients.
	//
	// It can only serve a single TCP client, the first one that
	// will connect to it.
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Print("-> ", string(netData))
		c.Write([]byte(netData))

		// It stops when it receives the STOP string as input.
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exitiong TCP server!")
			return
		}
	}
}
