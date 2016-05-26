package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	os.Setenv("WEB", "coolshell.cn")
	fmt.Println(os.Getenv("WEB"))

	for _, env := range os.Environ() {
		e := strings.Split(env, "=")
		println(e[0],"=",e[1])
	}

}
