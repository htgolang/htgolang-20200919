package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Pflag() string {
	var (
		dir string
	)
	flag.StringVar(&dir, "dir", "", "Please input director..")

	flag.Parse()

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	if dir == "" {
		flag.Usage()
		os.Exit(0)
	}
	return dir
}

func OpenFile(dirpath string) *os.File {
	file, err := os.Open(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func FDir(dirpath string, fileInfo os.FileInfo) {
	file, err := os.Open(dirpath + "/" + fileInfo.Name())
	// fmt.Println(dirpath, fileInfo.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileInfos, err := file.Readdir(-1)
	for _, fileinfo := range fileInfos {
		apath := filepath.Join(dirpath, fileInfo.Name(), fileinfo.Name())
		fileds := filepath.Ext(apath)
		// fmt.Println(fileds)
		if ok, _ := filepath.Match("*.go", fileds); ok {
			fmt.Println(filepath.Join(dirpath, fileInfo.Name(), fileinfo.Name()))
		}
	}
}

func Readdir(dirpath string) {
	file := OpenFile(dirpath)
	fileInfos, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() == true {
			FDir(dirpath, fileInfo)
		} else {
			fileds := filepath.Ext(fileInfo.Name())
			if ok, _ := filepath.Match("*.go", fileds); ok {
				fmt.Println(filepath.Join(dirpath, fileInfo.Name()))
			}
		}
	}
}

func main() {
	Readdir(Pflag())
}
