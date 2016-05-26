package main

import "fmt"

func main() {
	s1 := make([]int, 10)
	s1[1] = 100
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 5, 10)

	s2[4] = 200
	fmt.Println(s2, len(s2), cap(s2))
}
