package main

import "fmt"

type User struct{
	Id int
	Name string
}
type Manager struct {
	Group string
	User
}
type Tester interface {
	Test()
}
func (this *User)Test(){
	 fmt.Println(this)
}
func(this *Manager)Test(){
	fmt.Println(this)
}
func dosomething(o Tester){
	o.Test()
}
func main(){
	m:=Manager{User:User{2,"Jack"},Group:"IT"}
	dosomething(&m)
	dosomething(&m.User)
}
