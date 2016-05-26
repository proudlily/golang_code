package main

import "fmt"

type Tester interface {
	Test()
}
type MyTest struct{

}
func (this*MyTest) Test(){
	fmt.Println("Test")
}
type User struct {
	Id int
	Name string
	Tester
}
func main(){
	u:=User{100,"Jack",nil}
	u.Tester=new(MyTest)

	fmt.Println(u)
	u.Test()
}
