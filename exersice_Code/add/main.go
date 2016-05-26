package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"io/ioutil"
)

func main() {
	path := GetCurrPath()
	dir, _ := os.Open(path)
	files, _ := dir.Readdir(0)
	dir.Close()
	//遍历
	var data []byte
	var file *os.File
	for _, f := range files {
		if f.Name() == "main" || f.Name() == "all" || f.IsDir() == true {
			continue
		}
		file, _ = os.Open(path+"/"+f.Name())
		d, _ := ioutil.ReadAll(file)
		file.Close()
		data = append(data, d...)
	}

	file, _ = os.Create(path+"/all")
	file.Write(data)
	file.Close()
}
/*获取当前文件执行的路径*/

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	i := strings.LastIndex(path, "/")
	path = path[0:i]
	return path
}
