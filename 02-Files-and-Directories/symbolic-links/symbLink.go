package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Symbolic links are pointers to files or directories, which are
// resolved at the time of access. Symbolic links, which are also
// called soft links, are not equal to the file or the directory
// they are pointing to and are allowed to point to nowhere, which
// can sometimes complicate things.
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}
	filename := arguments[1]

	fileInfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println(filename, "is a symbolic link")
		realPath, err := filepath.EvalSymlinks(filename)
		if err == nil {
			fmt.Println("Path: ", realPath)
		}
	}
}
