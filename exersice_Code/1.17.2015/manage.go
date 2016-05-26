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
func main(){
	u:=User{1,"Tom"}
	m:=Manager{User:User{2,"Jack"},Group:"IT"}

	var i Tester

	i=&u
	i.Test()
	i=&m
	i.Test()
}
