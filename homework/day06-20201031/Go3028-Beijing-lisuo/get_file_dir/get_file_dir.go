package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var fileOrdir string

func main() {
	// cmd flag
	fileOrdir := flag.String("f", "file-or-dir", "specify the file or dir")
	flag.Parse()

	// get file'a abs path
	fileAbs, err := filepath.Abs(*fileOrdir)
	if err != nil {
		log.Fatal(err)
	}

	// stat file, and get os.FileInfo
	fileInfo, err := os.Stat(fileAbs)
	if os.IsNotExist(err) {
		log.Fatal(err)
	} else if fileInfo.IsDir() {
		// if a dir, read it and get a os.[]FileInfo
		fileInfoList, err := ioutil.ReadDir(fileAbs)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range fileInfoList {
			fmt.Printf("file: %-25v   dir: %v\n", file.Name(), file.IsDir())
		}

	} else {
		fileDir := filepath.Dir(fileAbs)
		fileInforList, err := ioutil.ReadDir(fileDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range fileInforList {
			fmt.Printf("file: %-25v   dir: %v\n", file.Name(), file.IsDir())
		}
	}

}
