package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"io/ioutil"
	"path/filepath"
)

func lookForAndMakeDir(toDir, date string) string {

	t, err := time.Parse("2006/1/2", date)
	if err != nil {
		t, err = time.Parse("2006/01/02", date)
		if err != nil {
			fmt.Println("日付ではありません")
			os.Exit(1)
		}
	}

	y := strconv.Itoa(t.Year())
	m := fmt.Sprintf("%02d", int(t.Month()))
	d := fmt.Sprintf("%02d", t.Day())

	targetDir := y + "/" + m + "/" + d

	_, err = os.Stat(toDir + targetDir)
	if err != nil {
		err = os.MkdirAll(toDir + targetDir, 0775)
		if err != nil {
			fmt.Println("ディレクトリの作成に失敗しました")
			os.Exit(1)
		}
	} else {
		fmt.Println("すでにあるディレクトリは選択できません")
		os.Exit(91)
	}

	return toDir + targetDir + "/"

}

func getFileList(fromDir string) []string {

	var files []string

	dir, err := ioutil.ReadDir(fromDir)
	if err != nil {
		fmt.Println("ディレクトリを開けませんでした")
		os.Exit(1)
	}

	for _, file := range dir {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	return files

}

func moveAndRenameFiles(fromDir, toDir, date, name string) {

	moveToDir := lookForAndMakeDir(toDir, date)

	files := getFileList(fromDir)

	for index, file := range files {
		num := fmt.Sprintf("%02d", index)
		ext := filepath.Ext(file)
		newName := name + "-" + num + ext
		err := os.Rename(fromDir + file, moveToDir + newName)
		if err != nil {
			fmt.Println("ファイルの移動に失敗しました")
			os.Exit(1)
		}
	}
}

func main() {

	fromDir := ""
	toDir := ""

	moveAndRenameFiles(fromDir, toDir, os.Args[1], os.Args[2])
}
