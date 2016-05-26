package main

import (
	"fmt"
	"os"
)

type File struct {
	nil string
}
type Error struct {
	inl string
}

func main() {

}

func Create(name string) (file *File, err Error) {
	os.Creat("astaxie.txt", 0666)
	if err != nil {
		fmt.Println(err)
	}
	return
}
