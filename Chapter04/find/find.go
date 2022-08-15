package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	printAll = false
	minusS   *bool // This is for printing socket files
	minusP   *bool // This is for printing pipes
	minusSL  *bool // This is for printing symbolic links
	minusD   *bool // This is for printing directories
	minusF   *bool // This is for printing files
)

func walkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if printAll {
		fmt.Println(path)
		return nil
	}

	mode := fileInfo.Mode()
	if mode.IsRegular() && *minusF {
		fmt.Println(path)
		return nil
	}

	if mode.IsDir() && *minusD {
		fmt.Println(path)
		return nil
	}

	fileInfo, _ = os.Lstat(path)

	if fileInfo.Mode()&os.ModeSymlink != 0 {
		if *minusSL {
			fmt.Println(path)
			return nil
		}
	}

	if fileInfo.Mode()&os.ModeNamedPipe != 0 {
		if *minusP {
			fmt.Println(path)
			return nil
		}
	}

	if fileInfo.Mode()&os.ModeSocket != 0 {
		if *minusS {
			fmt.Println(path)
			return nil
		}
	}

	return nil
}

func main() {
	minusS = flag.Bool("s", false, "Sockets")
	minusP = flag.Bool("p", false, "Pipes")
	minusSL = flag.Bool("sl", false, "Symbolic Links")
	minusD = flag.Bool("d", false, "Directories")
	minusF = flag.Bool("f", false, "Files")

	flag.Parse()
	flags := flag.Args()

	if *minusS && *minusP && *minusSL && *minusD && *minusF {
		printAll = true
	}

	if !(*minusS || *minusP || *minusSL || *minusD || *minusF) {
		printAll = true
	}

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	path := flags[0]

	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
