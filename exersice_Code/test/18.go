package main

import (
	"fmt"
	"runtime"
)

func test() {
	ps := make([]uintptr, 10)
	count := runtime.Callers(0, ps)

	for i := 0; i < count; i++ {
		f := runtime.FuncForPC(ps[i])
		fmt.Printf("%d,%s\n", i, f.Name())
	}
}
func main() {
	a := func() {
		test()
	}
	a()
}
