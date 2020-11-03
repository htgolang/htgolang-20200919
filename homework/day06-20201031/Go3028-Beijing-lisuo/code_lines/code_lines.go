package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var fileOrdir string
var tempFileList []string
var line int = 0

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
				absfile = filepath.Join(fileAbs, file.Name())
				data, err := ioutil.ReadFile(absfile)
				if err != nil {
					log.Fatal(err)
				}
				filedata := string(data)
				tempFile := strings.Split(filedata, "\n")
				tempFileList = append(tempFileList, tempFile...)
				//for range tempFile {
				//	line++
				//}
				//fmt.Printf("file %v has %v lines\n", file.Name(), line)
				//line = 0
			}
		}
		for range tempFileList {
			line++
		}
		fmt.Printf("all the files in dir %v end with .go&.cgo has %v lines\n", fileAbs, line)

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
				data, err := ioutil.ReadFile(absfile)
				//line := 0
				if err != nil {
					log.Fatal(err)
				}
				filedata := string(data)
				tempFile := strings.Split(filedata, "\n")
				tempFileList = append(tempFileList, tempFile...)
				//for range tempFile {
				//	line++
				//}
				//fmt.Printf("file %v has %v lines\n", file.Name(), line)
			}

		}
		for range tempFileList {
			line++
		}
		fmt.Printf("all the files at %v end with .go&.cgo has %v lines\n", fileDir, line)
	}

}
