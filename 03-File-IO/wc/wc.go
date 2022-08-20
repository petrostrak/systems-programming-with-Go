package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func countLines(filename string) (int, int, int) {
	var err error
	var numberOfLines int = 0
	var numberOfChars int = 0
	var numberOfWords int = 0

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", filename)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			os.Exit(1)
		}
		numberOfLines++
		for range regexp.MustCompile(`[^\\s]+`).FindAllString(line, -1) {
			numberOfWords++
		}
		numberOfChars += len(line)
	}

	return numberOfLines, numberOfWords, numberOfChars
}

func main() {
	minusC := flag.Bool("c", false, "Characters")
	minusW := flag.Bool("w", false, "Words")
	minusL := flag.Bool("l", false, "Lines")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Printf("usage: wc <file1> [<file2> [... <fileN]]\n")
		os.Exit(1)
	}

	totalLines := 0
	totalWords := 0
	totalChars := 0
	printAll := false

	for _, filename := range flag.Args() {
		numOfLines, numOfWords, numOfChars := countLines(filename)
		totalLines += numOfLines
		totalWords += numOfWords
		totalChars += numOfChars

		if (*minusC && *minusW && *minusL) || (!*minusC && !*minusW && !*minusL) {
			fmt.Printf("%d", numOfLines)
			fmt.Printf("\t%d", numOfWords)
			fmt.Printf("\t%d", numOfChars)
			fmt.Printf("\t%s\n", filename)
			printAll = true
			continue
		}

		if *minusL {
			fmt.Printf("%d", numOfLines)
		}
		if *minusW {
			fmt.Printf("\t%d", numOfWords)
		}
		if *minusC {
			fmt.Printf("\t%d", numOfChars)
		}
		fmt.Printf("\t%s\n", filename)
	}

	if (len(flag.Args()) != 1) && printAll {
		fmt.Printf("%d", totalLines)
		fmt.Printf("\t%d", totalWords)
		fmt.Printf("\t%d", totalChars)
		fmt.Println("\ttotal")
		return
	}

	if (len(flag.Args()) != 1) && *minusL {
		fmt.Printf("%d", totalLines)
	}
	if (len(flag.Args()) != 1) && *minusW {
		fmt.Printf("\t%d", totalWords)
	}
	if (len(flag.Args()) != 1) && *minusC {
		fmt.Printf("\t%d", totalChars)
	}
	if len(flag.Args()) != 1 {
		fmt.Printf("\ttotal\n")
	}
}
