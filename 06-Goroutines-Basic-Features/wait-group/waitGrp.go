package main

import (
	"fmt"
	"sync"
)

var (
	count int = 10
)

func main() {
	fmt.Println("Waiting for Goroutines!")

	var wg sync.WaitGroup

	// 	Calling sync.Add() before the Go statement
	// in order to prevent race conditions is important.
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	wg.Wait()
	fmt.Println("\nExiting...")
}
