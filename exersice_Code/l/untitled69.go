package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开文件，并进行相关处理
	f, err := os.Open("test.txt", os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	//文件关闭
	defer f.Close()

	//将文件作为一个io.Reader对象进行buffered I/O操作
	br := bufio.NewReader(f)
	for {
		//每次读取一行
		line, err := br.ReadString('\n')
		if err == os.EOF {
			break
		} else {
			fmt.Printf("%v", line)
		}
	}
}
