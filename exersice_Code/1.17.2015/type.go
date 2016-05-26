package main

import "fmt"

type User struct{
	Id   int
	Name string
}

func Test(i interface{}) {
	switch v := i.(type){
	case *User:
		println("User:Name=", v.Name)
	case string:
		println("string=", v)
	default:
		fmt.Printf("%T:%v\n", v, v)

	}
}
func main() {
	Test(&User{1, "Tom"})
	Test("Hello World!")
	Test(123)
}

