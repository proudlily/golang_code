package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte("I am a beautiful women :-) ")
	for a := 0; a < 2; a++ {
		h := md5.New()
		h.Write(data)
		data = []byte(hex.EncodeToString(h.Sum(nil)))
	}
	fmt.Printf("%s\n", data)
}
