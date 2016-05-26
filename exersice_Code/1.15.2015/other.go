package main

func arg()(int,string){
	return 1,"abc"
}

func test(i int,s string){
	println(i,s)
}

func main(){
	test(arg())
}

