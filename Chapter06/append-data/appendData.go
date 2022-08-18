package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s message filename\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	message := os.Args[1]
	filename := os.Args[2]

	// O_RDWR opens the file read-write.
	// O_APPEND appends data to the file when writing.
	// O_CREATE creates a new file if none exists.
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", message)
}
