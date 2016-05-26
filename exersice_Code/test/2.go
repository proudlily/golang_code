package main

import "fmt"

func main() {

	l := func(s string) int {
		fmt.Println("get length")
		return len(s)
	}
	ss := "ilovelily"

	for i := 0; i < l(ss); i++ {
		fmt.Println(ss[i])
	}
	fmt.Println("-------")

	for i, m := 0, l(ss); i < m; i++ {
		fmt.Println(ss[i])
	}

}
