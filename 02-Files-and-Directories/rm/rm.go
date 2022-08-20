package main

import (
	"fmt"
	"os"
)

// The rm.go file is a Go implementation of the rm(1) tool that
// illustrates how you can delete files in Go.
func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	file := arguments[1]
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
		return
	}
}
