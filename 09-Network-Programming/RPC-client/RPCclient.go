package main

import (
	"fmt"
	RPCserver "go-systems-programming/09-Network-Programming/RPC-server"
	"net/rpc"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string!")
		os.Exit(100)
	}
	CONNECT := arguments[1]

	c, err := rpc.Dial("tpc", CONNECT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args := RPCserver.MyInts{17, 18, true, false}
	var reply int

	err = c.Call("MyInterface.Add", args, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Reply (Add): %d\n", reply)

	err = c.Call("MyInterface.Subtract", args, &reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	fmt.Printf("Reply (Subtract): %d\n", reply)
}
