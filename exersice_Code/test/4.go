package main

import "fmt"

func main() {
	switch a := 5; {

	case a > 1:
		fmt.Println("a")
	case a > 2:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}
