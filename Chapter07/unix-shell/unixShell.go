package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	VERSION string = "0.2"
)

func main() {
	// Present prompt
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		// Read a line
		line := scanner.Text()
		words := strings.Split(line, " ")

		// Get the first word of the line
		command := words[0]

		// If it is a built-in shell command, execute the command
		switch command {
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		case "version":
			fmt.Println(VERSION)
		default:
			// otherwise, echo the command
			fmt.Println(line)
		}

		fmt.Print("> ")
	}

}
