package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"time"
)
//习题三,tail -f 实现
func Tailf() error {
	//读取参数
	args := os.Args
	if len(args) != 2 {
		return fmt.Errorf("请输入一个参数作为查看的对象!")
	}
	path := args[1]
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	//判断状态以及是否为文件夹，对于文件夹抛出异常
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("不能以文件夹作为查看对象!\n")
	}
	//将光标seek到文件末尾
	file.Seek(0, 2)
	//用bufio打开文件
	reader := bufio.NewReader(file)
	for {
		//读取内容,如果报错为文件末尾则等待0.1s,如果报错非文件末尾直接返回报错
		text, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Print(string(text))
	}
	return nil
}

func main() {
	err := Tailf()
	if err != nil {
		fmt.Println(err)
	}
}