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

	// get abs path
	fileAbs, err := filepath.Abs(*fileOrdir)
	if err != nil {
		log.Fatal(err)
	}

	// stat the file
	fileInfo, err := os.Stat(fileAbs)
	if os.IsNotExist(err) {
		log.Fatal(err)
	} else if fileInfo.IsDir() {
		// if a dir, read it and get a []FileInfo
		fileInfoList, err := ioutil.ReadDir(fileAbs)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range fileInfoList {
			// each file's abs path
			absfile := filepath.Join(fileAbs, file.Name())
			suffix := filepath.Ext(absfile)
			if suffix == ".go" || suffix == ".cgo" {
				fmt.Printf("file: %-25v   dir: %v\n", file.Name(), file.IsDir())
			}
		}

	} else {
		fileDir := filepath.Dir(fileAbs)
		fileInforList, err := ioutil.ReadDir(fileDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range fileInforList {
			absfile := filepath.Join(fileDir, file.Name())
			suffix := filepath.Ext(absfile)
			if suffix == ".go" || suffix == ".cgo" {
				fmt.Printf("file: %-25v   dir: %v\n", file.Name(), file.IsDir())
			}

		}
	}

}
