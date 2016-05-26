package main

import (
	"fmt"
	"time"

)


var over =make (chan bool)
var sem=make (chan int ,2)


func worker(i int){
	sem <-1
	fmt.Println(time.Local.Format("04:45"),i)

	time.Sleep(1*time.Second)

	<- sem

	if i==9{
		over <- true
	}
}

func main(){
   for i:=0;i<10;i++{
	   go worker(i)
   }

  <- over
}
