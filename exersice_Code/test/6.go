package main

import "fmt"

func main() {
	var a struct{ x, y int }
	var b struct{ x, y int }
	a.x = 1
	a.y = 2
	b = a
	fmt.Println(b)
}
