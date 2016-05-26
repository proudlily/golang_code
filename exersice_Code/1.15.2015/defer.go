package main

func main(){
	x:=10
	defer  func(a int)(){
         println("a=",a)
	}(x)

	defer println("print=",x)

	x+=100
	println(x)
}
