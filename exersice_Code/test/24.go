package main

import (
	"fmt"
)

func main() {

	s1 := make([]int, 3, 6)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println(s2, len(s2), cap(s2))

	s1 = s1[:cap(s1)]

	fmt.Println(s1, len(s1), cap(s1))

	s3 := append(s2, []int{4, 5, 6}...)

	s3[0] = 100

	fmt.Println(s3, len(s3), cap(s3))

	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println(s2, len(s2), cap(s2))
}
