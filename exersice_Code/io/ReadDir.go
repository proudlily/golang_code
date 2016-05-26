package main

import (
	"fmt"
	"io/ioutil"
	//"strings"
)

func main() {
	rd, err := ioutil.ReadDir("/home/lily")
	for _, fi := range rd {
		fmt.Println("")
		fmt.Println(fi.Name())
		fmt.Println(fi.IsDir())
		fmt.Println(fi.Size())
		fmt.Println(fi.ModTime())
		fmt.Println(fi.Mode())
	}
	fmt.Println("")
	fmt.Println(err)
}
