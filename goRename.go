package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
)

func fileList() []string {

	dir, _ := os.Getwd()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		os.Exit(1)
	}

	var items []string

	for _, file := range files {
		if !file.IsDir() {
			items = append(items, file.Name())
		}
	}

	return items
}

func rename(name string, files []string) {

	var num, ext, newName string

	for index, file := range files {
		ext = filepath.Ext(file) 
		num = fmt.Sprintf("%02d", index + 1)
		newName = fmt.Sprintf("%s-%s%s", name, num, ext)
		err := os.Rename(file, newName)
		if err != nil {
			os.Exit(1)
		}
	}
}

func main() {
	
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println("引数が足りないか多すぎます")
		fmt.Println("使用法:goRename　変更したいファイル名")
		os.Exit(1)
	}

	files := fileList()

	rename(os.Args[1], files)
}
