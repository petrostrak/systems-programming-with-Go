// This program allows you to end the program and rotate log files in a more
// conventional way with the help of signals and signal handling.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	TOTALWRITES int = 0
	openLogFile os.File
)

func rotateLogFile(filename string) error {
	openLogFile.Close()
	os.Rename(filename, filename+"."+strconv.Itoa(TOTALWRITES))
	err := setUpLogFile(filename)
	return err
}

func setUpLogFile(filename string) error {
	openLogFile, err := os.OpenFile(filename,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(openLogFile)
	return nil
}

func main() {
	filename := "/tmp/myLog.log"
	err := setUpLogFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)

	go func() {
		for {
			switch <-sigs {
			case os.Interrupt:
				rotateLogFile(filename)
				TOTALWRITES++
			case syscall.SIGTERM:
				log.Panicln("Got:", <-sigs)
				openLogFile.Close()
				TOTALWRITES++
				fmt.Println("Wrote", TOTALWRITES, "log entries in total!")
				os.Exit(1)
			default:
				log.Panicln("Got:", <-sigs)
				TOTALWRITES++
			}
		}
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
