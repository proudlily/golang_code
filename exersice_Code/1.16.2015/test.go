package main

func test()(int,string){
	return 123,"baby"
}
func main(){
	_,a:=test()
	println(a)//_,a输出的是baby
}
