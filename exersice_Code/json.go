package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name   string
	Age    string
	Height string
}
type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{Name: "weidao", Age: "12", Height: "180"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
