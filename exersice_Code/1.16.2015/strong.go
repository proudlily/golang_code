package main

import "fmt"

type User struct {
	Id   int
	Name string
}
type Tester interface {
	Test()
}

func (this User) Test() {

}
func main(){
	u:=User{100,"Tom"}

	i:=Tester(u)
	u.Id=100
	u.Name="Jack"
	fmt.Println(u,i)

	//fmt.Println(&(i.(User)))
}
