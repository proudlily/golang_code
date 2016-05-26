package main

import "fmt"


type User struct {
	Id int
	Name string
}
func main(){
    i:=100

	var p *int =&i
	println(*p)

	up:=&User{1,"jack"}
	up.Id=100
	fmt.Println(up)

    u2 :=*up
	u2.Name="Tom"
	fmt.Println(u2,up)
}
