package main

type callback func(s string)

func test(a,b int,sum func(int,int) int ){
	println(sum(a,b))
}
func main(){
	var cb callback
	cb=func (s string){
		println(s)
	}
	cb("Hello World!")

	test(1,2,func(a ,b int)int{return a+b})
}
