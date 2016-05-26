package main

import (
	"runtime"
	"time"
)

func main() {
	println("start:", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func(n int) {
			println(n, runtime.NumGoroutine())
		}(i)
		time.Sleep(3 * time.Second)
		println("over:", runtime.NumGoroutine())
	}

}
