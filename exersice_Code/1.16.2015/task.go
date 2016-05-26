package main

import "time"

func task(arg ...int) chan int {
	x := 0
	c := make(chan int)
	go func() {

		time.Sleep(2 * time.Second)

		for _, v := range arg {
			x+=v

		}
		c <-x
	}()
	return c

}
func main(){
	c:=task(1,2,3,4)

	println("do something....")

	i:=<- c
	println("task result=",i)

}
