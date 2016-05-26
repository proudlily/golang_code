package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {
	var jack string
	i := 100
	var p *int = &i
	fmt.Println(*p)

	up := &User{1, jack}
	up.Id = 100
	fmt.Println(up)

	u2 := *up
	u2.Name = "tom"
	fmt.Println(up, u2)
}
