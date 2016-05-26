package main

import "fmt"

func test(a, b int) (c int) {
	defer func() {

		c += 100
	}()

	c = a + b
	return
}

func main() {
	x := test(10, 0)
	fmt.Println(x)
}
