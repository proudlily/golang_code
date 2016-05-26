package main

func main(){
	i:=0
	p:=&i

	i++
	a:=i

	*p++

	b:=*p

	println(a,b)
}
