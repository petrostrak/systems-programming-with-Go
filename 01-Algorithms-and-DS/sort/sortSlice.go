package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	a := aStructure{"Petros", 175, 65}

	mySlice = append(mySlice, a)
	a = aStructure{"Maggie", 165, 55}
	mySlice = append(mySlice, a)
	a = aStructure{"Maria", 160, 65}
	mySlice = append(mySlice, a)
	a = aStructure{"George", 180, 90}
	mySlice = append(mySlice, a)

	fmt.Println("0: ", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight < mySlice[j].weight
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println(">:", mySlice)

}
