package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// CopyFile 拷贝文件函数
func copyFile(srcFile, dstFile string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", srcFile, err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", dstFile, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //调用io.Copy()拷贝内容
}

func main() {
	//定义命令行参数方式1
	var srcFile string
	var dstFile string
	flag.StringVar(&srcFile, "s", "", "源文件")
	flag.StringVar(&dstFile, "d", "", "目的文件")

	//解析命令行参数
	flag.Parse()

	_, err := copyFile(srcFile, dstFile)
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")
}
