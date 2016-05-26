package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP string
}
type Serverslice struct {
	Servers []Server
}
func main(){
	var s Serverslice
	str := `{"Servers":
	[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},
	{"ServerName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)  //反序列化
	fmt.Println(s)
}
