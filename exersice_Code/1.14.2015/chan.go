package main

func recv(c <-chan int,over chan <-bool){
	for v :=range c{
		println(v)
	}
	over <- true
}

func send(c chan <-int){
	for i:=0;i<10;i++{
	c <- i
}
	close (c)
}

func main(){
	c :=make (chan int)
	o:=make (chan bool)

	go recv(c,o)
	go send(c)

	<- o
}

