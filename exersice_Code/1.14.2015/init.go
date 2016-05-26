package main

import (
	"strings"
	"time"
)

func init(){
	println(a)
	println("init()is over")
}

var a string =strings.Join([]string{"a","b"},",")

func init(){
	time.Sleep(2 *time.Second)
	println("init2 over...")
}
func main(){
	println("hello world!")
}
