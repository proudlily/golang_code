package main

import (
	"fmt"

)

type Node struct {
	_     int
	Value string
	Nest *Node
}

func main() {
	a := &Node{Value:"a"}
	b := &Node{Value:"b", Nest:a}

	fmt.Printf("%#v/n", a)
	fmt.Printf("%#v/n", b)
}
