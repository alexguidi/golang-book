package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	createFile()
	readFile()
	showContentFolder()
	showContentFolderAndSubfolder()
}

func readFile() {

	fmt.Println("-----Reading a file-----")
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	fmt.Println("-----Reading a file in a short way-----")
	//short way to do the same thing above
	bss, errr := ioutil.ReadFile("test.txt")
	if errr != nil {
		return
	}
	strs := string(bss)
	fmt.Println(strs)
}

func createFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString("Hi my name is Alex")
}

func showContentFolder() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func showContentFolderAndSubfolder() {
	filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
