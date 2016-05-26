package main

import (
	"bytes"
	"fmt"
)

func main(){
	s:=bytes.Buffer{}

	s.WriteString("Hello")
	s.WriteString("World")
	s.WriteString("!")

	fmt.Println(s.String())
}
