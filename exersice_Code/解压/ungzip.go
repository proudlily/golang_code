package main

import (
	"fmt"
	//"io"
	//	"os"
	// "time"
	"archive/tar"
	//"compress/gzip"
	"archive/zip"
)

func main() {
	// file read
	fr, err := os.Open("/home/lily/Work/video-js-4.5.1.zip")
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	// gzip read
	gr, err := zip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	// tar read
	tr := tar.NewReader(gr)

	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// 显示文件
		fmt.Println(h.Name)

		// 打开文件
		fw, err := os.OpenFile("file2/"+h.Name, os.O_CREATE|os.O_WRONLY, 0644 /*os.FileMode(h.Mode)*/)
		if err != nil {
			panic(err)
		}
		defer fw.Close()

		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			panic(err)
		}

	}

	fmt.Println("un tar.gz ok")
}
