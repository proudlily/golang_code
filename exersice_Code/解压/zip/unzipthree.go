package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func unzipFile(f *zip.File, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	path := filepath.Join(dest, f.Name)
	if f.FileInfo().IsDir() {
		err := os.MkdirAll(path, f.Mode())
		if err != nil {
			return err
		}
	} else {
		f, err := os.OpenFile(
			path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}
	return nil
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		err := unzipFile(f, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := Unzip("/home/lily/Work/aaa.zip", "/home/lily/Work")
	if err != nil {
		log.Fatal(err)
	}
}
