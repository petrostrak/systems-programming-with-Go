package main

import (
	"fmt"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)

	go func() {
		defer close(temp)
		w.Wait()
	}()

	select {
	// When the appropriate number of w.Done() calls have been
	// executed, the 1st case of the select will be executed.
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(1)

	t := 2 * time.Second

	if timeout(&w, t) {
		fmt.Println("Timed out")
	} else {
		fmt.Println("OK!")
	}

	w.Done()

	if timeout(&w, t) {
		fmt.Println("Timed out")
	} else {
		fmt.Println("OK!")
	}
}
