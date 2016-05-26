package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

/*const (
	LOGFILEPATH = "/home/lily/Work/go/src/company/upload/case/unzip.log"
)*/

func main() {
	/*logfile, err := os.OpenFile(LOGFILEPATH, os.O_CREATE|os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	if logger == nil {
		fmt.Println("logger init error")
	}*/
	r, err := zip.OpenReader("/home/lily/Work/video-js-4.5.1.zip")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Println("FileName : ", f.Name)
		rc, err := f.Open()
		if err != nil {
			logger.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68) //打印文件内容
		if err != nil {
			if err != io.EOF {
				logger.Fatal(err)
			}
		}
	}
}
