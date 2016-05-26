package main

import "strings"

func main(){
	a,b,c:="a","b","c"
	s:=strings.Join([]string{a,b,c},"")
	println(s)
}
