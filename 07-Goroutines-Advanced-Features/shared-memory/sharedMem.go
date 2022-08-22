// In Go terminology, this is sharing by communicating
// instead of communicating by sharing
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	readValue  = make(chan int)
	writeValue = make(chan int)
)

func setValue(newValue int) {
	writeValue <- newValue
}

func getValue() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	go monitor()
	var wg sync.WaitGroup

	for r := 0; r < 20; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			setValue(rand.Intn(100))
		}()
	}

	wg.Wait()
	fmt.Printf("\nLast value: %d\n", getValue())
}
