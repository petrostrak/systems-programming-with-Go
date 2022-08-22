// A mutex works like a buffered channel of capacity 1 that allows at
// most one goroutine to access a shared variable at a time. This
// means that there is no way for two or more goroutines to try to
// update that variable simultaneously.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var (
	mu             sync.Mutex
	sharedVariable string = ""
)

// A critical section cannot be embedded in another critical section
// when both critical sections use the same Mutex variable
func addDot() {
	mu.Lock()
	sharedVariable += "."
	mu.Unlock()
}

func read() string {
	mu.Lock()
	a := sharedVariable
	mu.Unlock()
	return a
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s n\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	numGR, _ := strconv.ParseInt(os.Args[1], 10, 64)

	var wg sync.WaitGroup
	var i int64
	for i = 0; i < numGR; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addDot()
		}()
	}

	wg.Wait()
	fmt.Printf("-> %s\n", read())
	fmt.Printf("Length: %d\n", len(read()))
}
