// Read a JSON record and convert it into one Go variable. (decode)
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

func fromJSON(filename string, key any) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}

	return nil
}

// go run readJSON.go ../save-JSON-data/record
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		os.Exit(100)
	}
	filename := arguments[1]

	var myRecord Record
	err := fromJSON(filename, &myRecord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myRecord)
	}
}
