package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Value int
}

var root = new(*Tree)

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{Left: nil, Right: nil, Value: v}
	}

	if v == t.Value {
		return t
	}

	if v > t.Value {
		t.Left = insert(t.Left, v)
		return t
	}

	t.Right = insert(t.Right, v)
	return t
}

func traverse(t *Tree) {
	if t == nil {
		return
	}

	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n)
		t = insert(t, temp)
	}
	return t
}

func main() {
	tree := create(30)
	traverse(tree)
	fmt.Println()
	fmt.Println("The value of the root is: ", tree.Value)
}
