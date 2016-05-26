package main

import (
	"fmt"

)

func main(){
	s:="abc"
	var c byte=s[1]

	fmt.Printf("%c,%02x\n",c,c)
	bs:=[]byte(s)
	bs[1]='B'
	println(string(bs))
}
