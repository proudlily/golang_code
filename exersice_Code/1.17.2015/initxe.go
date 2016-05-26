package main

import "time"

func init(){
	go func(){
		println("goroutine is init....")
		time.Sleep(1*time.Second)
	}()
	println("init is over ....")
}

func main(){
	println("Hello World...")
	time.Sleep(2*time.Second)// 等待 goroutines 执行行结束后才退出,否则看不到效果了。
}
