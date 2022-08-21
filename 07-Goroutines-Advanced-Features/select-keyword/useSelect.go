package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func createNumber(max int, randomNumberChan chan<- int, done chan bool) {
	for {
		// The select statement waits for a channel
		// to unblock and then executes on that.
		select {
		case randomNumberChan <- rand.Intn(max):
		case x := <-done:
			if x {
				close(done)
				close(randomNumberChan)
				return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	randomNumberChan := make(chan int)
	done := make(chan bool)

	if len(os.Args) != 3 {
		fmt.Printf("usage: %s count max\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	n1, _ := strconv.ParseInt(os.Args[1], 10, 64)
	count := int(n1)
	n2, _ := strconv.ParseInt(os.Args[2], 10, 64)
	max := int(n2)

	fmt.Printf("Going to create %d random numbers.\n", count)

	go createNumber(max, randomNumberChan, done)
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", <-randomNumberChan)
	}

	done <- false
	fmt.Println()

	_, ok := <-randomNumberChan
	if ok {
		fmt.Println("Chann3el is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	done <- true
	_, ok = <-randomNumberChan
	if ok {
		fmt.Println("Chann3el is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
