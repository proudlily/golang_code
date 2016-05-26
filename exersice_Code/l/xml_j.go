package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:"innerxml"`
}
type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")   //打开口号中名字的文件
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(v)

}
