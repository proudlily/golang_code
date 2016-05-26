package main

import (
	"crypto/md5"
	"encoding/hex"

	"fmt"
)

func main() {
	h := md5.New()
	var A []byte
	A = h.Write([]byte("I am a gentle men "))
	for a := 0; a < 1; a++ {

		h.Write([]byte(A))
	}
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))
}
