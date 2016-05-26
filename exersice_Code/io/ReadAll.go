package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	ra, _ := ioutil.ReadFile("/home/lily/文档/额外文件")
	fmt.Printf("%s", ra)
}
