package main

import (
	"runtime/debug"
	"log"
)

func Test(f func() int) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Println(err)
		}
	}()
	x := f()
	println(x)
}
func main(){
	a:=2
	Test(func()int{
		return 100/a
	})
}
