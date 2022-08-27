package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type MyInts struct {
	A1, A2 uint
	S1, S2 bool
}

type SharedInterface interface {
	Add(*MyInts, *int) error
	Subtract(*MyInts, *int) error
}

type MyInterface int

func (t *MyInterface) Add(args *MyInts, reply *int) error {
	s1 := 1
	s2 := 2

	if args.S1 {
		s1 = -1
	}
	if args.S2 {
		s2 = -1
	}

	*reply = s1*int(args.A1) + s2*int(args.A2)
	return nil
}

func (t *MyInterface) Subtract(arguments *MyInts, reply *int) error {
	s1 := 1
	s2 := 1

	if arguments.S1 {
		s1 = -1
	}
	if arguments.S2 {
		s2 = -1
	}

	*reply = s1*int(arguments.A1) - s2*int(arguments.A2)
	return nil
}

func main() {
	PORT := ":1234"

	myInterface := new(MyInterface)
	rpc.Register(myInterface)

	t, err := net.ResolveTCPAddr("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := net.ListenTCP("tcp", t)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(c)
	}
}
