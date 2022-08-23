package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// go run -race rd.go 10
func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Printf("usage: %s number\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}
	numGR, _ := strconv.ParseInt(os.Args[1], 10, 64)

	var wg sync.WaitGroup
	var i int64

	for i = 0; i < numGR; i++ {
		wg.Add(1)

		// As the anonymous function takes no arguments, it uses
		// the current value of i , which cannot be determined
		// with any certainty as it depends on the operating system
		// and the Go scheduler.
		go func() {
			defer wg.Done()
			fmt.Printf("%d ", i)
		}()
	}

	wg.Wait()
	fmt.Println("\nExiting...")
}
