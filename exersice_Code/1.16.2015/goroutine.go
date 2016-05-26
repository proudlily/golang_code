package main

import "time"

func main(){
	go func(){
     time.Sleep(3*time.Second)
		println("go....")
	}()
	println("hi!")
	time.Sleep(5*time.Second)
}
