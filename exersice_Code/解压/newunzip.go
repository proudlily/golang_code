package main

import (
	"archive/zip"
	"fmt"
	//"io"
	//"os"
)

func main() {
	r, err := zip.OpenReader("/home/lily/Work/video-js-4.5.1.zip")
	if err != nil {
		panic(err)
	}

	for _, f := range r.File {
		fmt.Println("FileName : ", f.Name)
		rc, err := f.Open()
		if err != nil {
			panic(err)
		}
		/* _, err = io.CopyN(os.Stdout, rc, 68) //打印文件内容
		if err != nil {
			if err != io.EOF {
				logger.Fatal(err)
			}
		}*/
	}
}
