package main
import (
	"html/template"
	"os"
	"fmt"
)

type Person struct {
	UserName string
}

func main(){
	t := template.New("")  //创建一个模板
	t, _ = t.Parse("hello {{.UserName}}!")  //解析
	p := Person{UserName:"Astaxie"}  //获得信息？
	t.Execute(os.Stdout, p)  //执行模板操作
}
