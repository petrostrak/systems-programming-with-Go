// This is a program that uses a Unix socket, which is a special
// Unix file, to read and write data.
package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

func readSocket(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print("Read: ", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("unix", "/tmp/aSocket.sock")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go readSocket(c)
	n := 0

	for {
		message := []byte("Hi there: " + strconv.Itoa(n) + "\n")
		c.Write(message)
		time.Sleep(5 * time.Second)
		n += 1
	}
}
