package main

import "fmt"

func closures(x int) func(int) {

	return func(y int) int {
		return x + y
	}

}

func main() {
	f := closures(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

}
