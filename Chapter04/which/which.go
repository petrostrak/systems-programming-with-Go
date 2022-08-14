package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// The which(1) utility searches the value of the PATH environment
// variable in order to find out if an executable file can be found
// in one of the directories of the PATH variable.
func main() {
	minusA := flag.Bool("a", false, "a")
	flag.Parse()

	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	file := flags[0]

	foundIt := false

	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")

	for _, dir := range pathSlice {
		fullPath := dir + "/" + file

		// The call to os.Stat() tells whether the file we are
		// looking for actually exists or not.
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()

			// The mode.IsRegular() function checks whether the file is
			// a regular file or not because we are not looking for
			// directories or symbolic links
			if mode.IsRegular() {

				// Below statement verifies that the file is actually an
				// executable file using a binary operation.
				if mode&0111 != 0 {
					foundIt = true
					if *minusA {
						os.Exit(0)
					} else {
						fmt.Println(fullPath)
						os.Exit(0)
					}
				}
			}
		}
	}
	if !foundIt {
		os.Exit(1)
	}
}
