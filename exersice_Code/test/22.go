package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func main() {

	a := [...]User{
		{0, "User0"},
		{1, "User1"},
	}

	b := [...]*User{

		{0, "User0"},
		{1, "User1"},
	}

	fmt.Println(a)
	fmt.Println(b, b[1])
}
