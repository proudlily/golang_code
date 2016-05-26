package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Email    string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!{{.Email}}")
	p := Person{UserName: "Astaxie", Email: "annqlm@qq.com"}
	t.Execute(os.Stdout, p)
}
