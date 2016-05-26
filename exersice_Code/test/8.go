package main

import "fmt"

func test2() (x int, y string) {
	x = 2
	y = "test2"
	return

}

func test1() (x int, y string) {

	x = 1
	y = "test1"

	m, y := test2()
	fmt.Println(m, y)

	fmt.Println(y)
	return
}

func main() {
	_, _ = test1()

}
