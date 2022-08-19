// A small Go program that lists all the processes of a Unix
// machine by executing a Unix command and getting its output
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	ps, err := exec.LookPath("ps")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ps)

	command := []string{"ps", "-a", "-x"}
	env := os.Environ()
	err = syscall.Exec(ps, command, env)
	if err != nil {
		fmt.Println(err)
	}
}
