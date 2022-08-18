package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

var (
	MAX int = 90
	MIN int = 0
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s length\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	LENGTH, _ := strconv.ParseInt(os.Args[1], 10, 64)

	f, err := os.Open("/dev/random")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	rand.Seed(seed)
	f.Close()

	fmt.Println("Seed:", seed)

	startChar := "!"
	var i int64

	for i = 0; i < LENGTH; i++ {
		anInt := int(random(MIN, MAX))
		newChar := string(startChar[0] + byte(anInt))
		if newChar == " " {
			i -= i
			continue
		}
		fmt.Print(newChar)
	}
	fmt.Println()
}
