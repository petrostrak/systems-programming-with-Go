// A channel is a communication mechanism that allows goroutines to exchange data.
// However, some rules exist here.
//
// First, each channel allows the exchange of a particular data type, which is also
// called the element type of the channel, and
// second, for a channel to operate properly, you will need to use some Go code to
// receive what is sent via the channel.
// Additionally, as each channel has its own type, the developer should define it.
// Last, when you are using a channel as a function parameter, you
// can specify its direction, that is, whether it will be used for writing or reading.
package main

import (
	"fmt"
	"time"
)

func writeChan(c chan<- int, x int) {
	fmt.Println(x)
	c <- x // blocking because nobody is reading from the c channel
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan<- int)
	go writeChan(c, 10)
	time.Sleep(2 * time.Second)
}
