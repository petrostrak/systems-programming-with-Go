// Both io.Writer and io.Reader are interfaces that
// embed the io.Write() and io.Read() methods, respectively.
//
// The program computes the characters of its input file
// and writes the number of characters to another file—if
// you are dealing with Unicode characters that take more
// than one byte per character, you might consider that
// the program is reading bytes. The output filename has
// the name of the original file plus the .Count extension.
package main

import (
	"fmt"
	"io"
	"os"
)

func countChars(r io.Reader) int {
	buf := make([]byte, 16)
	total := 0
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}
		if err == io.EOF {
			break
		}
		total += n
	}

	return total
}

func writeNumberOfChars(w io.Writer, x int) {
	fmt.Fprintf(w, "%d\n", x)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]

	_, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("error on file %s: %s\n", filename, err)
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s: %s\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	chars := countChars(f)
	filename = filename + ".Count"
	f, err = os.Create(filename)
	if err != nil {
		fmt.Printf("error creating file %s: %s\n", filename, err)
		os.Exit(1)
	}
	defer f.Close()
	writeNumberOfChars(f, chars)
}
