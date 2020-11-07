package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 获取相对路径的信息(执行目录)
	fmt.Println(".")
	fmt.Println(filepath.Abs("."))
	fmt.Println(os.Getwd())

	// 二进制程序所在的路径信息
	fmt.Println(os.Args[0])
	fmt.Println(filepath.Abs(os.Args[0]))
	path, _ := filepath.Abs(os.Args[0])
	fmt.Println(filepath.Dir(path))
}
