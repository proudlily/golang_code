package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`    //标签，使只能输出大写的能输出小写
	ServerIP string `json:"serverIP,omitempty"`// 如果 ServerIP 为空,则不输出到 JSON 串中
}
type Serverslice struct {
	Servers []Server
}
func main(){
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName:"Shanghai_VPN",ServerIP:"127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName:"Beijing_VPN",ServerIP:"127.0.0.2"})
	b, _ := json.Marshal(s)

	fmt.Println(string(b))
}
