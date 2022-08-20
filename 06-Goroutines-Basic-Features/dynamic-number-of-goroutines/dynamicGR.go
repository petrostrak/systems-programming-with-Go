package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s integer\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	numGR, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	var i int64
	wg.Add(int(numGR))

	for i = 0; i < numGR; i++ {
		go func(x int64) {
			defer wg.Done()
			fmt.Printf(" %d", x)
		}(i)
	}
	wg.Wait()
	fmt.Println("\nExiting...")
}
