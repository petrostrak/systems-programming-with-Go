// Buffered channels allow the Go scheduler to put jobs in the queue
// quickly in order to be able to serve more requests. Moreover, you
// can use buffered channels as semaphores in order to limit throughput.
// The technique works as follows: incoming requests are forwarded to
// a channel, which processes one request at a time. When the channel
// is done, it sends a message to the original caller saying that it
// is ready to process a new request. So, the capacity of the buffer
// of the channel restricts the number of simultaneous requests it can
// keep and process
package main

import "fmt"

func main() {
	// You can write five integers to that channel without having
	// to read any one of them in order to make space for the others.
	numbers := make(chan int, 5)

	counter := 10
	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter*2; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done.")
		}
	}
}
