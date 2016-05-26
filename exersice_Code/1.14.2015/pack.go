package main

import "fmt"

type MyInt int
type MySlice []int

func (i MyInt) test() {
	fmt.Println("myInt=%v/n", i)

}
func (i MySlice) test() {
	fmt.Println("mySlice=%v/n", i)

}

func main() {
	i:=MyInt(10)
	i.test()

	s := MySlice{100, 200}
	s.test()
}
