package main

import (
	"fmt"
	"runtime"
)

func test() {
	fmt.Println(runtime.Caller(1))
}

func main() {
	test()
}
