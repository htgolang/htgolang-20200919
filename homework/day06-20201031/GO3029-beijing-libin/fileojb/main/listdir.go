package main

import (
	"fileobj/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	d1 string = "tmp"
	d2 string = "/private/tmp"
)
var dirlist []string = []string{}

func main() {
	// dirlist = make([]string, 1)

	ok := utils.CheckDirExist(d2)
	if !ok {
		fmt.Printf("%s 不存在\n", d2)
	}

	err := filepath.Walk(d2, DirRecur)
	if err != nil {
		fmt.Println("walk err", err.Error())
	}
}

func DirRecur(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		// fmt.Println(filepath.Base(path))
		//与.go或者.cgo字符串比较
		if strings.HasSuffix(filepath.Base(path), ".go") || strings.HasSuffix(filepath.Base(path), ".cgo") {
			lines, _ := utils.GetFileLines(path)
			fmt.Printf("文件是%s,%s的行数是%d\n", filepath.Base(path), filepath.Base(path), lines)
		}
	} else {
		dirlist = append(dirlist, path)
	}
	if path == "" {
		return err
	}
	return nil
}
