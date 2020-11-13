package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs(os.Args[1])

	_, err := filename(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allfile)

	gofile := make([]string, 0, 10)
	cgofile := make([]string, 0, 10)
	for _, name := range allfile {
		shuffix := filepath.Ext(name)
		if shuffix == ".go" {
			gofile = append(gofile, name)
		} else if shuffix == ".cgo" {
			cgofile = append(cgofile, name)
		}
	}
}

var allfile []string = make([]string, 0, 20)

func filename(path string) ([]string, error) {
	var filenames = make([]string, 0, 10)

	fileinfo, err := os.Stat(path)
	if err != nil {
		return nil, err //如果文件路径不存在则报错
	}

	if fileinfo.IsDir() {
		//目录
		finfos, _ := ioutil.ReadDir(path)
		for _, info := range finfos {
			if info.IsDir() {
				filename(filepath.Join(path, info.Name()))
			} else {
				filenames = append(filenames, info.Name())
			}
		}
		// fmt.Println(path)
		// fmt.Printf("\t%v\n\n", filenames)
		allfile = append(allfile, filenames...)
		return filenames, nil
	}
	//文件
	return []string{fileinfo.Name()}, nil

}
