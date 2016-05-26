package main

import "fmt"

func sum(s string, args ...int) {
	var x int
	for _, n := range args {
		x += n
	}
	fmt.Println(s, x)
}

func main() {
	sum("1+2+3=", 1, 2, 3)

	x := []int{0, 1, 2, 3, 4}
	sum("0+1+2=", x[:3]...)
}
