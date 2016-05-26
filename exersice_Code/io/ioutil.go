package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//func ReadAll(r io.Reader) ([]byte, error)

func main() {
	s := strings.NewReader("Hello World!")
	ra, _ := ioutil.ReadAll(s)
	fmt.Printf("%s", ra)
	// Hello World!
}
