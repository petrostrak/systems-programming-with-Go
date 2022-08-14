package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	file := arguments[1]

	// The FileInfo structure describes the file or directory
	// examined by os.Stat().
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// From the FileInfo structure, you can discover the permissions
	// of a file by calling the Mode() function.
	mode := info.Mode()
	fmt.Print(file, ": ", mode, "\n")
}
