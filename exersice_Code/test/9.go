package main

import "fmt"

func A(c chan int) {

	c <- 10

}
func main() {

	c := make(chan int)

	go A(c)

	fmt.Println(<-c)

}
