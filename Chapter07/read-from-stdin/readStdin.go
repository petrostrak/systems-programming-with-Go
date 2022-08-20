// Unix pipes in Go
//
// Pipes have two serious limitations: first, they usually communicate in one direction,
// and second, they can only be used between processes that have a common ancestor. The
// general idea behind pipes is that if you do not have a file to process, you should wait
// to get your input from standard input. Similarly, if you are not told to save your
// output to a file, you should write your output to standard output, either for the user
// to see it or for another program to process it. As a result, pipes can be used for
// streaming data between two processes without creating any temporary files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := ""
	var f *os.File
	arguments := os.Args

	// Resolve whether there is an actual file to process
	//
	// if not, read data from stdin.
	if len(arguments) == 1 {
		f = os.Stdin // os.Stdin implements io.Reader
	} else {
		filename = arguments[1]
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening %s: %s", filename, err)
			os.Exit(1)
		}
		f = fileHandler
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
