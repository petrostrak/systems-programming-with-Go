// Log files tend to get bigger and bigger all the time because
// data is written to them all the time; it would be good to
// have a technique for rotating them.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	TOTALWRITES       int = 0
	ENTRIESPERLOGFILE int = 100
	WHENTOSTOP            = 230
	openLogFile       os.File
)

func rotateLogFile(filename string) error {
	openLogFile.Close()
	os.Rename(filename, filename+"."+strconv.Itoa(TOTALWRITES))
	err := setUpLogFile(filename)
	return err
}

func setUpLogFile(filename string) error {
	openLogFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(openLogFile)
	return nil
}

func main() {
	numberOfLogEntries := 0
	filename := "/tmp/myLog.log"
	err := setUpLogFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for {
		log.Println(numberOfLogEntries, "This is a test log entry")
		numberOfLogEntries++
		TOTALWRITES++

		if numberOfLogEntries > ENTRIESPERLOGFILE {
			rotateLogFile(filename)
			numberOfLogEntries = 0
		}
		if TOTALWRITES > WHENTOSTOP {
			rotateLogFile(filename)
			break
		}
	}
	fmt.Println("Wrote", TOTALWRITES, "log entries!")
}
