package main

import "fmt"

func main(){
	s:="a:中国人"
	println("uft-8 string",s,len(s))

	bs:=[]byte(s)
	bs[0]='A'
	for i:=0;i<len(bs);i++{
		fmt.Printf("%02x",bs[i])
	}
	println(string(bs))

	u:=[]rune(s)
	println("unicode len=",len(u))

	for i:=0;i<len(u);i++{
		fmt.Printf("%04x",u[i])
	}

	u[4]='龙'
	println(string(u))
}
