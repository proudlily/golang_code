package main

import (
	"unsafe"
	"fmt"
)

const N int = int(unsafe.Sizeof(0))

func main() {
	x := 0x1234
	p := unsafe.Pointer(&x)
	p2 := (*[N]byte)(p)

	for i, m := 0, len(p2); i < m; i++ {
		fmt.Println("%02x", p2[i])
	}


}
