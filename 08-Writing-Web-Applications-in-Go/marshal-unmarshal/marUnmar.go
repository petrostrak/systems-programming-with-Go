package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := Record{
		Name:    "Petros",
		Surname: "Trak",
		Tel: []Telephone{
			{Mobile: true, Number: "1234-abcd"},
			{Mobile: false, Number: "abcc-567"},
		},
	}

	// equivalent to json.encode but returns a []byte.
	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(rec))

	var record Record

	// equivalent to json.decode
	err = json.Unmarshal(rec, &record)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(record)
	}
}
