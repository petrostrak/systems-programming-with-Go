package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var (
	printAll = false
	minusS   *bool   // This is for printing socket files
	minusP   *bool   // This is for printing pipes
	minusSL  *bool   // This is for printing symbolic links
	minusD   *bool   // This is for printing directories
	minusF   *bool   // This is for printing files
	minusX   *string // This is for excluding files
	minusRE  *string // This is for printing files from regex
)

func walkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if excludenames(path, *minusX) {
		return nil
	}

	if !regularExpression(path, *minusRE) {
		return nil
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

func excludenames(name, exclude string) bool {
	if exclude == "" {
		return false
	}
	if filepath.Base(name) == exclude {
		return true
	}
	return false
}

func regularExpression(name, regExp string) bool {
	if regExp == "" {
		return true
	}
	r, err := regexp.Compile(regExp)
	if err != nil {
		return false
	}
	return r.MatchString(regExp)
}

func main() {
	minusS = flag.Bool("s", false, "Sockets")
	minusP = flag.Bool("p", false, "Pipes")
	minusSL = flag.Bool("sl", false, "Symbolic Links")
	minusD = flag.Bool("d", false, "Directories")
	minusF = flag.Bool("f", false, "Files")
	minusX = flag.String("x", "", "files")

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
