package main

import (
	"runtime"
	"time"
)
func main(){
	go func(){
	   defer println("go1 defer...")

		for i:=1;i<10;i++{
			println("go1:",i)
			if i==5{
				runtime.Goexit()
			}
		}

	}()

	go func(){
	   println("go2")
	}()

	time.Sleep(3*time.Second)
}
