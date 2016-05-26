package main

import ("os"
	"fmt"
	"log"


)

func main(){
	logfile,err:=os.OpenFile("/home/lily/Work/go/src/exersice_Code/1.2.015/map.go",os.O_RDWR|os.O_CREATE,0)
    if err!=nil {
        fmt.Println("%s\r\n",err.Error())
		os.Exit(-1);
	}
	defer logfile.Close()
	logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("hello")
	logger.Fatal("test")
	logger.Fatal("test2")


}

