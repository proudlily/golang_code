package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	k := template.New("fieldname example")
	k, _ = k.Parse(`hello {{.UserName}}!
{{range .Emails}}
an email {{.}}
{{end}}
{{with .Friends}}
{{range .}}
my friend name is {{.Fname}}
{{end}}
{{end}}
`)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	k.Execute(os.Stdout, p)
}
