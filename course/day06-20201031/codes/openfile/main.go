package main

import (
	"fmt"
	"os"
	"time"
)

/*
O_RDONLY 以只读方式打开
O_WRONLY 以只写方式打开
O_RDWR  以读写方式打开
O_APPEND 追加
O_CREATE 文件不存在则创建
O_EXCL   文件必须不存在
// O_SYNC   使用同步IO
O_TRUNC  截断文件
*/

// os.Open => 读，文件不存在报错
func Open(name string) (*File, error) {
	return os.OpenFile(name, os.O_RDONLY, 0777)
}

// os.Create => 写文件，文件存在 截断 文件不存在 创建
func Create(name string) (*File, error) {
	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
}

func main() {
	// os.Open => 读，文件不存在报错
	// os.Create
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%s\n", time.Now().Format("2006-01-02 15:04:05"))
}
