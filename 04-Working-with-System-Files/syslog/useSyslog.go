package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])

	// Returns a writer that tells your program where to direct
	// all log messages.
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	}
	sysLog.Crit("Crit: Logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_ALERT|syslog.LOG_LOCAL7, "Some program!")
	if err != nil {
		log.Fatal(sysLog)
	}
	sysLog.Emerg("Emerg: Logging on Go!")

	fmt.Fprintf(sysLog, "log.Print: Logging in Go!")
}
