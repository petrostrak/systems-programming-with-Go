// A Go program that handles the SIGTERM and
// SIGINT signals.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("Got", sig)
}

func main() {
	// The definition of a channel, which acts as a way of passing data around, that is
	// required for the technique ( sigs ).
	sigs := make(chan os.Signal, 1)

	// 	Calling signal.Notify() in order to define the list of signals you want to be
	// able to catch.
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	// 	The definition of an anonymous function that runs in a goroutine right after
	//  signal.Notify() , which is used for deciding what you are going to do when you
	//  get any of the desired signals.
	go func() {
		for {
			sig := <-sigs
			fmt.Println(sig)
			handleSignal(sig)
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(10 * time.Second)
	}
}
