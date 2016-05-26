package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) string() string {
	return "<<" + h.name + "-" + strconv.Itoa(h.age) + "years" + h.phone + ">>"
}
func main() {
	Bob := Human{"Bob", 39, "000"}
	fmt.Println("This Human is :", Bob)
}
