// A signal channel is a channel that is used just for signaling
// The program executes four goroutines—when the first one is
// finished, it sends a signal to a signal channel by closing it,
// which will unblock the second goroutine. When the second
// goroutine finishes its job, it closes another channel that
// unblocks the remaining two goroutines.
package main

import (
	"fmt"
	"time"
)

// The A() function is blocked by the
// channel defined in the a parameter.
func A(a, b chan struct{}) {
	<-a
	fmt.Println("A!")
	time.Sleep(time.Second)
	close(b)
}

// Similarly, the B() function is blocked by the channel
// stored in the b argument, which means that until the b
// channel is closed, the B() function will be waiting in
// its first statement.
func B(b, c chan struct{}) {
	<-b
	fmt.Println("B!")
	close(c)
}

// Once again, the C() function is blocked by the channel
// stored in its a argument.
func C(a chan struct{}) {
	<-a
	fmt.Println("C!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	// Here, you start four goroutines. However, until you
	// closes the a channel, all of them will be blocked!
	//
	// A() will finish first and unblock B() that will unblock the two C()
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)
	close(x)
	time.Sleep(2 * time.Second)
}
