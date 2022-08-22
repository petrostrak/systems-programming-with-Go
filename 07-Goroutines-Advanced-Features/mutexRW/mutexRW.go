// Go offers another type of mutex, called sync.RWMutex , that
// allows multiple readers to hold the lock but only a single
// writer - sync.RWMutex is an extension of sync.Mutex that adds
// two methods named sync.RLock and sync.RUnlock , which are used
// for locking and unlocking for reading purposes. Locking and
// unlocking a sync.RWMutex for exclusive writing should be done
// with Lock() and Unlock() , respectively.
package main

import (
	"fmt"
	"sync"
	"time"
)

type secret struct {
	sync.RWMutex
	counter  int
	password string
}

var (
	Password = secret{counter: 1, password: "password"}
)

// This function makes changes to one of its arguments, which means
// that it requires an exclusive lock, hence the use of the Lock()
// and Unlock() functions.
func change(c *secret, pass string) {
	c.Lock()
	fmt.Println("LChange")
	time.Sleep(20 * time.Second)
	c.counter += 1
	c.password = pass
	c.Unlock()
}

func show(c *secret) string {
	fmt.Println("LShow")
	time.Sleep(time.Second)

	c.RLock()
	defer c.RUnlock()
	return c.password
}

func counts(c secret) int {
	c.RLock()
	defer c.RUnlock()
	return c.counter
}

func main() {
	fmt.Println("Pass:", show(&Password))

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Go Pass:", show(&Password))
		}()
	}

	go func() {
		change(&Password, "123456")
	}()

	fmt.Println("Pass:", show(&Password))
	time.Sleep(time.Second)
	fmt.Println("Counter:", counts(Password))
}
