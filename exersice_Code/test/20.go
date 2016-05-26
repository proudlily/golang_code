package main

import "fmt"

func test(a *[10]int) {
	a[2] = 100
}

func main() {
	var a = new([10]int)
	test(a)
	fmt.Println(a, len(a))
}
