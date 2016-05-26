package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)
func main() {
	data := []byte("123456")
       h := md5.New()
		h.Write(data)
		data = []byte(hex.EncodeToString(h.Sum(nil)))
fmt.Printf("%s\n", data)
}
