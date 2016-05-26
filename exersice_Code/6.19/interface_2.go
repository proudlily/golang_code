package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) string() string {
	return "<<" + h.name + "-" + strconv.Itoa(h.age) + "years" + h.phone + ">>"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = Person{"Dennis", 70}
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d]is an int and its value is %d\n", index, value)
		}
	}
}
