package main

import "time"

func test(c chan int) {
	time.Sleep(3 * time.Second)
	println("go....")
	c <-1
}
func main(){
	c :=make (chan int)
	go test(c)

	println("hi...")
	<-c
	println("over....")
}
