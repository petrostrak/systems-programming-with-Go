// Nil channels are a special sort of channel
// that will always block.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The addIntegers() function stops after the time
// defined in the time.NewTimer()
func addIntegers(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C:
			fmt.Println(sum)
		}
	}
}

// The sendIntegers() function keeps generating random
// numbers and sends them to the c channel as long as
// the c channel is open.
func sendIntegers(c chan int) {
	for {
		c <- rand.Intn(100)
	}
}

func main() {
	c := make(chan int)

	go addIntegers(c)
	go sendIntegers(c)
	time.Sleep(2 * time.Second)
}
