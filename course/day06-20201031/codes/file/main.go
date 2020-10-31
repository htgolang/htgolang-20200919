package main

import (
	"fmt"
	"os"
)

func main() {
	// path := "test.txt"
	path := "testdir"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// open file dir 都可以打开
	// file => 读取文件内容
	// dir => 读取目录下的文件名
	// 如何判断是文件 还是 目录
	fileInfo, err := file.Stat()
	fmt.Println(fileInfo, err)

	fmt.Println(fileInfo.Name())  // 文件名
	fmt.Println(fileInfo.Size())  // 目录 => 0
	fmt.Println(fileInfo.IsDir()) // 是否为文件夹
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())
	fmt.Println(int(fileInfo.Mode()))

	// fmt.Println(file.Readdirnames(-1))
	fileInfos, err := file.Readdir(-1)

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name(), fileInfo.IsDir())
	}
	// dir => 如何读取目录下的文件名
}
