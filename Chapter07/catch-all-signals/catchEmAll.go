package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("* Got:", sig)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handleSignal(sig)
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(1)
			case syscall.SIGUSR1:
				handleSignal(sig)
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(10 * time.Second)
	}
}
