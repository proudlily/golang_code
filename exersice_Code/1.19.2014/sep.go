package main

import (
	//"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	f, _ := os.Open("sep.go")
	//fmt.Println(v)
	defer f.Close()

	r := bufio.NewReaderSize(f, 4096)

	for {
		line, isPrefix, err := r.ReadLine()

		if isPrefix {
			print(string(line))// 如果是一一行行的前部分,则不换行行。(行行数据超出缓冲区大大小小)

		}else if len(line) > 0 {
			println(string(line))
		}
		if err == io.EOF {
			break
		}

	}


}
