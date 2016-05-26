package main

import (
	"log"
	"os"
	"encoding/binary"
	"fmt"
)

func checkError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func write() {
	f, err := os.Create("test.dat")
	checkError(err)

	defer func() {
		f.Sync()// 必须显式调用用, Close 不会刷新。
		f.Close()
	}()

	var i int32 = 0x123 // 必须是具体⻓长度的类型
	checkError(binary.Write(f, binary.LittleEndian, i))

	var d float32 = 0.1234
	checkError(binary.Write(f, binary.LittleEndian, d))

	var s string = "Hello China"
	checkError(binary.Write(f, binary.LittleEndian, int32(len(s))))// 先写入入字符串⻓长度的类型
	_, err = f.WriteString(s)// 写入入字符串内容
	//fmt.Println(err)
	checkError(err)
}
func read() {
	f, err := os.Open("test.dat")
	checkError(err)

	defer f.Close()

	var i int32
	checkError(binary.Read(f, binary.LittleEndian, &i))

	var d float32
	checkError(binary.Read(f, binary.LittleEndian, &d))

	var l int32
	checkError(binary.Read(f, binary.LittleEndian, &l))

	s := make([]byte, l)
	_, err = f.Read(s)

	fmt.Printf("%#x;%f;%s;\n", i, d, string(s))


}
func main() {
	write()
	read()
}
