package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	path        string
	newPath     string
	permissions fs.FileMode
	minusTEST   *bool
)

func walkFunction(currentPath string, info os.FileInfo, err error) error {
	fileInfo, _ := os.Lstat(currentPath)
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println("Skipping", currentPath)
		return nil
	}

	fileInfo, err = os.Stat(currentPath)
	if err != nil {
		fmt.Println("*", err)
		return err
	}

	mode := fileInfo.Mode()
	if mode.IsDir() {
		tempPath := strings.Replace(currentPath, path, "", 1)
		pathToCreate := newPath + "/" + filepath.Base(path) + tempPath

		if *minusTEST {
			fmt.Println(":", pathToCreate)
			return nil
		}

		_, err := os.Stat(pathToCreate)
		if os.IsNotExist(err) {
			os.MkdirAll(pathToCreate, permissions)
		} else {
			fmt.Println("Did not create", pathToCreate, ":", err)
		}
	}

	return nil
}

func main() {
	minusTEST = flag.Bool("test", false, "Test run!")
	flag.Parse()

	flags := flag.Args()
	if len(flags) == 0 || len(flags) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	path = flags[0]
	newPath = flags[1]
	permissions = os.ModePerm

	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		os.MkdirAll(newPath, permissions)
	} else {
		fmt.Println(newPath, "already exists - quitting...")
		os.Exit(1)
	}

	err = filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
