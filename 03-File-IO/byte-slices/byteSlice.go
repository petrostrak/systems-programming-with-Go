package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename!")
		os.Exit(1)
	}
	filename := os.Args[1]

	aByteSlice := []byte("Petros Trakadas!\n")
	os.WriteFile(filename, aByteSlice, 0644)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	anotherByteSlice := make([]byte, 100)
	n, err := f.Read(anotherByteSlice)
	if err == io.EOF {
		fmt.Println("EOF:", err)
	}
	fmt.Printf("Read %d bytes: %s", n, anotherByteSlice)
}
