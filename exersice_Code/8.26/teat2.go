package main

import "fmt"

type A struct {
	b string
	c int
}

func (a *A) B() {
	fmt.Println(a.b)
}
func main() {
	d := A{
		b: "flown",
	}
	d.B()

}
