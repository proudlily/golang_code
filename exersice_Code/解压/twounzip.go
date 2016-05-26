package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	reader, err := zip.OpenReader("/home/lily/Work/video-js-4.5.1.zip")
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	writer, err := os.Create("/home/lily/Work/video-js-4.5.1.unzip")
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
