package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunction(path string, info os.FileInfo, err error) error {
	filename, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := filename.Mode()
	if mode.IsDir() {
		fmt.Println(path)
	}

	return nil
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	path := arguments[1]

	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
