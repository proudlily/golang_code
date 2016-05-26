package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	//"path/filepath"
)

func main() {
	reader, err := zip.OpenReader("/home/lily/Work/aaa.zip")
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	writer, err := os.Create("/home/lily/Work/aaa.zip")
	if err != nil {
		panic(err)
	}

	defer writer.Close()

	for _, f := range reader.Reader.File {

		zipped, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err = io.Copy(writer, zipped); err != nil {
			panic(err)
		}

		zipped.Close()
		fmt.Println("File uncompressed\n")
	}

}
