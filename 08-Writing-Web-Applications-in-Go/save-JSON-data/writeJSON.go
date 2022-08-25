// Convert regular data into JSON format in order to
// send it over a network connection. (encode)
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

func toJSON(filename string, key any) {
	out, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	encodeJSON := json.NewEncoder(out)
	err = encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// go run writeJSON.go record
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		os.Exit(100)
	}
	filename := arguments[1]

	myRecord := Record{
		Name:    "Petros",
		Surname: "Trak",
		Tel: []Telephone{
			{Mobile: true, Number: "1234-abcd"},
			{Mobile: false, Number: "abcc-567"},
		},
	}

	toJSON(filename, myRecord)
}
