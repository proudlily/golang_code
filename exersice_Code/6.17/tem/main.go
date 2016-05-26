package main

import (
	"log"
	"net/http"
	"text/template"
	"fmt"
)

func A(w http.ResponseWriter, r *http.Request) {
	t := template.New("A.tmpl") //创建一个模板
	data := map[string]string{}
	data["miao"]="miao"
	t, err := t.ParseFiles("A.tmpl") //解析模板文件
	if err != nil{
		fmt.Println(err)
	}
	err =t.Execute(w,data) //执行模板写操作
	if err != nil{
		fmt.Println(err)
	}
}

func B(w http.ResponseWriter, r *http.Request) {
	t := template.New("B.tmpl") //创建一个模板
	data := map[string]string{}
	data["miao"]="miao"
	t, err := t.ParseFiles("B.tmpl") //解析模板文件
	if err != nil{
		fmt.Println(err)
	}
	err =t.Execute(w,data) //执行模板写操作
	if err != nil{
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/A", A)       //设置访问的路由
	http.HandleFunc("/B", B)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
