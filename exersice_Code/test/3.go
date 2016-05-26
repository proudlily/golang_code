package main

import "fmt"

func main() {
	switch a := 5; a {
	case 0, 1:
		fmt.Println("a")
	case 100:
	case 5:
		fmt.Println("b")
		fallthrough
	default:
		fmt.Println("c")

	}

}
