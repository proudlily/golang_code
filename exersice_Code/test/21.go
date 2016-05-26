package main

import "fmt"

func main() {
	var a = [3][2]int{[...]int{1, 2}, [...]int{3, 4}}
	var b = [3][2]int{{1, 2}, {3, 4}}

	c := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	c[1][1] = 100

	fmt.Println(a, "\n", b, "\n", c, len(c), len(c[0]))
}
