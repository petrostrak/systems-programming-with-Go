// The program adds a line number in front of each line of a text
// file. This means that you will need to read the input file line
// by line, keep a variable that will hold the line number value,
// and save it using the original name.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	minusINIT := flag.Int("init", 1, "Initial value")
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("usage: insertLineNumber <files>\n")
		os.Exit(1)
	}

	lineNumber := *minusINIT
	for _, filename := range flags {
		fmt.Println("Processing:", filename)
		r, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lines := strings.Split(string(r), "\n")

		for i, line := range lines {
			lines[i] = fmt.Sprintf("%d: %s", i, line)
			lineNumber++
		}

		lines[len(lines)-1] = ""
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println("Processed", lineNumber-*minusINIT, "lines!")
}
