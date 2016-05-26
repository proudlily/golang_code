package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Name   string
	Age    string
	Height string
}
type Serverslice struct {
	Servers []Server
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	var s Serverslice
	s.Servers = append(s.Servers, Server{Name: "weidao", Age: "12", Height: "180"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))

	r.ParseForm() //解析参数,默认是不会解析的
	//fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//fmt.Println("key:", k)
	//fmt.Println("val:", strings.Join(v, ""))
	//}
	//fmt.Fprintf(w, "Hello as taxie!")
	//这个写入到 w 的是输出到客户端的
}
func main() {
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
