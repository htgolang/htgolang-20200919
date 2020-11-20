package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"bufio"
	"io"
	"sync"
)
//作业四:递归打印文件夹下所有以.go和.cgo结尾的文件名称，并统计代码行数
func IsGo(name string) bool {
	return strings.HasSuffix(name, ".go") || strings.HasSuffix(name, ".cgo")
}

func SumCount(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	file, _ := os.Open(path)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else if err == io.EOF {
			break
		}
		count++
	}
	fmt.Printf("%v中包含%v行代码\n", path , count)
}

func WalkAllGo(path string, wg *sync.WaitGroup) {
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
			WalkAllGo(path + "/" + fileInfo.Name(), wg)
		} else if !fileInfo.IsDir() && IsGo(fileInfo.Name()){
			wg.Add(1)
			SumCount(path + "/" + fileInfo.Name(), wg)
		}
	}
}

func main() {
	var wg sync.WaitGroup
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
	WalkAllGo(path, &wg)
	wg.Wait()
}

