// The GOMAXPROCS environment variable (and Go function) allows you
// to limit the number of operating system threads that can execute
// user-level Go code simultaneously.
package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())
}
