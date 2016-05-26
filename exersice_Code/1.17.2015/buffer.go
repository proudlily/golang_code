package main

import "fmt"

type User struct{
	Id   int
	Name string
}
type Tester interface{
	Test()
}
type Stringer interface{
	String()
}
func( this User) Test(){
	fmt.Println("Test:",this)
}
func (this*User) String(){
	fmt.Printf("String:Id=%d,Name=%s/n",this.Id,this.Name)
}
func main(){
	u:=User{1,"Tom"}
	p:=&u
	Tester(u).Test()
	Tester(p).Test()
	Stringer(p).String()
}
