package main

import "strings"

func init() {
	println(a)// 在全局变量初始化完成后才会执行行,否则初始化就有问题了......
	println("init is over....")
}

var a string = strings.Join([]string{"a", "b"}, "")

func init() {
	println("init2 is over...")// main 必须等待初始化完成后才被执行行,包括其他包中的初始化函数。
}

func main() {
	println("Hello World!")
}
