// Large files that are created with the os.Seek() function may have holes in them and
// occupy fewer disk blocks than files with the same size, but without holes in them; such files
// are called sparse files.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// go run . 100000 test
// dd if=/dev/urandom bs=1 count=100000 of=noSparseDD
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s SIZE filename\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	SIZE, _ := strconv.ParseInt(os.Args[1], 10, 64)
	filename := os.Args[2]

	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("File %s already exists.\n", filename)
		os.Exit(1)
	}

	fd, err := os.Create(filename)
	if err != nil {
		log.Fatal("Failed to create output")
	}

	_, err = fd.Seek(SIZE-1, 0)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to seek")
	}

	_, err = fd.Write([]byte{0})
	if err != nil {
		fmt.Println(err)
		log.Fatal("Write operation failed")
	}

	err = fd.Close()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Failed to close file")
	}
}
