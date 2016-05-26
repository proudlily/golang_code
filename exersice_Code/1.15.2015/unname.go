package main

import "fmt"


type MyInt int

type IntPointer *int

type MySlice []int

func test(a,b int ,c string){
	println(a,b,c)
}

func add(a,b int)int{
	return a+b
}
func main(){
	var x=123
	var p IntPointer=&x

	var a MySlice=[]int{1,2,3}
     var b []int=a

	u:=struct {
		name string
		age int
		}{"user1",12}

    test(1,2,"abc")
	println(add(1,2))


	fmt.Println(*p)
    fmt.Println(a,b)
	fmt.Println(u)
}

