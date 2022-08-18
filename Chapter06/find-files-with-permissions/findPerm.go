// Finding files based on their permissions
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var PERMISSIONS string

func permissionsOfFile(filename string) string {
	info, err := os.Stat(filename)
	if err != nil {
		return "-1"
	}

	return info.Mode().String()[1:10]
}

func walkFunction(path string, info os.FileInfo, err error) error {
	_, err = os.Lstat(path)
	if err != nil {
		return err
	}

	if permissionsOfFile(path) == PERMISSIONS {
		fmt.Println(path)
	}

	return err
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Printf("usage: %s RootDirectory permissions\n",
			filepath.Base(arguments[0]))
		os.Exit(1)
	}
	path := arguments[1]

	path, _ = filepath.EvalSymlinks(path)
	PERMISSIONS = arguments[2]

	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
	}
}
