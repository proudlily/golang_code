package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "asatxie.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

}
