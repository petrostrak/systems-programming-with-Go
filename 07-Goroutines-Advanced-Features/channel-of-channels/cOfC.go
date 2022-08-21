// Two possible reasons to use such a channel are as follows:
// For acknowledging that an operation finished its job
// For creating many worker processes that will be controlled
// by the same channel variable
package main

import "fmt"

var numbsers = []int{0, -1, 2, 3, -4, 5, 6, -7, 8, 9, 10}

// The f1() function returns integer numbers that belong to the numbers variable.
func f1(cc chan chan int, done chan struct{}) {
	c := make(chan int)
	cc <- c
	defer close(c)

	total := 0
	i := 0
	for {
		select {
		case c <- numbsers[i]:
			i += 1
			i %= len(numbsers)
			total += 1
		case <-done:
			c <- total
			return
		}
	}
}

func main() {
	c := make(chan chan int)
	f := make(chan struct{})

	go f1(c, f)
	data := <-c

	i := 0
	for integer := range data {
		fmt.Printf("%d ", integer)
		i += 1
		if i == 100 {
			close(f)
		}
	}
	fmt.Println()
}
