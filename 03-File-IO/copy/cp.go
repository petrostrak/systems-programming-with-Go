package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var BUFFERSIZE int64

func Copy(src, dst string, buffersize int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	buf := make([]byte, buffersize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf); err != nil {
			return err
		}
	}

	return err
}

// go run . ../io.Copy/copiedText test 2048
func main() {
	if len(os.Args) != 4 {
		fmt.Printf("usage: %s source destination BUFFERSIZE\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	source := os.Args[1]
	dest := os.Args[2]
	BUFFERSIZE, _ = strconv.ParseInt(os.Args[3], 10, 64)

	fmt.Printf("Copying %s to %s\n", source, dest)
	err := Copy(source, dest, BUFFERSIZE)
	if err != nil {
		fmt.Printf("File copying failed: %q\n", err)
	}
}
