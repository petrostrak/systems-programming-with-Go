// A pipeline is a method for connecting goroutines so that
// the output of a goroutine becomes the input of another
// with the help of channels. The benefits of using pipelines
// are as follows:
//
// One of the benefits you get from using pipelines is that
// there is a constant flow in your program because nobody
// waits for everything to be completed in order to start the
// execution of goroutines and channels of the program.
// Additionally, you are using less variables and therefore
// less memory space because you do not have to save everything
// Last, the use of pipelines simplifies the design of the
// program and improves its maintainability.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func genNumbers(min, max int64, out chan<- int64) {
	var i int64
	for i = min; i < max; i++ {
		out <- i
	}
	close(out)
}

func findSquares(out chan<- int64, in <-chan int64) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func calculateSum(in <-chan int64) {
	var sum int64 = 0
	for x := range in {
		sum += x
	}
	fmt.Printf("The sum of squares is %d\n", sum)
}

// 				  naturals				 squares
// 				  channel 				 channel
// genNumbers(...) ----> findSquares(...) ----> calculateSum(...)
// 	  ^												^
// 	  |												|
// 	  |												|
//   input										  output

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s n1 n2\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	n1, _ := strconv.ParseInt(os.Args[1], 10, 64)
	n2, _ := strconv.ParseInt(os.Args[2], 10, 64)
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		os.Exit(1)
	}

	naturals := make(chan int64)
	squares := make(chan int64)
	go genNumbers(n1, n2, naturals)
	go findSquares(squares, naturals)
	calculateSum(squares)
}
