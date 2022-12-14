package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}
	input := os.Args[1]

	buf, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	in := string(buf)
	s := bufio.NewScanner(strings.NewReader(in))
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		fmt.Print(s.Text())
	}
}
