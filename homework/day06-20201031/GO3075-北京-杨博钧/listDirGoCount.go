package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"bufio"
	"io"
)
//作业四:递归打印文件夹下所有以.go和.cgo结尾的文件名称，并统计代码行数
func IsGo(name string) bool {
	return strings.HasSuffix(name, ".go") || strings.HasSuffix(name, ".cgo")
}

func CountLines(path string) (int, error) {
	count := 0
	file, _ := os.Open(path)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return 0, err
		} else if err == io.EOF {
			break
		}
		count++
	}
	return count, nil
}

func WalkAllGo(path string) {
	//打开路径
	filePath, _ := os.Open(path)
	defer filePath.Close()
	fileInfos, err := filePath.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	//判断是否为文件夹，不是则递归
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			WalkAllGo(path + "/" + fileInfo.Name())
		} else if !fileInfo.IsDir() && IsGo(fileInfo.Name()){
			count, err := CountLines(path + "/" + fileInfo.Name())
			if err != nil {
				fmt.Printf("检测到%v但是无法解析代码行数,原因为:%v\n", path + "/" + fileInfo.Name(), err)
			}
			fmt.Printf("%v中包含%v行代码\n", path + "/" + fileInfo.Name(), count)
		}
	}
}

func main() {
	//获取参数作为遍历路径
	//判断输入的参数个数
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("请传一个参数作为扫描路径!\n")
	}
	path := args[1]
	//获取文件状态，并判断是否为文件夹
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	if !stat.IsDir() {
		log.Fatalf("输入的地址不是一个路径!\n")
	}
	WalkAllGo(path)
}
