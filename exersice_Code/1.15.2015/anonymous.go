package main

func closures(x int)(func(int)int){
	return func(y int) int {
      return x+y
	}
}
func main(){
	f:=closures(10)

	println(f(1))
	println(f(2))
}
