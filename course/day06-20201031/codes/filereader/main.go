package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	path := "test.txt"
	file, err := os.Open(path)
	fmt.Println(file, err)
	if err != nil {
		return
	}

	defer file.Close()

	content := make([]byte, 3)
	// make([]Type, len, cap)
	// make(p[])

	for {
		n, err := file.Read(content)
		if err != nil {
			// EOF => 标识文件读取结束了
			// 非EOF
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(content[:n]))
	}

}
