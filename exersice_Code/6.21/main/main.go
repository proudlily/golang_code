package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func A(w http.ResponseWriter, r *http.Request) {
	t := template.New("A.tmpl")

	data := map[string]string{}
	t, err := t.ParseFiles("A.tmpl")
	if err != nil {
		fmt.Println(err)

	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func B(w http.ResponseWriter, r *http.Request) {
	t := template.New("B.tmpl")
	data := map[string]string{}
	t, err := t.ParseFiles("B.tmpl")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/A", A)
	http.HandleFunc("/B", B)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
