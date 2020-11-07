package main

import (
	"fmt"
	"log"
	"os"
)
//作业二:递归打印文件夹下所有文件名称
func WalkAll(path string){
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
			WalkAll(path + "/" + fileInfo.Name())
		} else {
			fmt.Println(path + "/" + fileInfo.Name())
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
	WalkAll(path)
}
