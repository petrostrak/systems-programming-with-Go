// The technique in this section will use the ioutil.WriteFile() and ioutil.ReadFile()
// functions. Note that ioutil.ReadFile() does not implement the io.Reader interface
// and therefore is a little restrictive.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide source and destination filenames")
		os.Exit(1)
	}

	sourceFile := os.Args[1]
	destFile := os.Args[2]

	r, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Printf("could not read from source file: %s", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(destFile, r, 0664)
	if err != nil {
		fmt.Printf("could not write to destination file: %s", err)
		os.Exit(1)
	}
}
