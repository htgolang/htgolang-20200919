package main

import (
	"fmt"
	"io"
	"os"
)

const bufferSize = 1024

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: tailf path")
		return
	}
	path := os.Args[1]
	// path 不存在
	// path 存在 文件夹

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	content := make([]byte, bufferSize)
	for {
		n, err := file.Read(content)
		if err != nil {
			if err == io.EOF {
				// 等待
			} else {
				fmt.Println(err)
				break
			}
		} else {
			fmt.Print(string(content[:n]))
		}
	}

}
