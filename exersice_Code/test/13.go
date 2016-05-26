package main

import "fmt"

func test(a, b int) int {
	defer println("defer1:", a, "/", b)

	defer func() {
		println("defer2:", a, b)
	}()
	return a / b
}

func main() {
	a := test(10, 2)
	b := test(10, 0)

	fmt.Print(a, b)
}
