package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("123")

	content := make([]byte, 10)
	// 123
	// file.Seek(0, 0) => 123
	file.Seek(-2, 1)
	fmt.Println(file.Read(content))
	fmt.Println(content)

	file.Seek(0, 0)

	file.WriteString("ab") // 文件内 ab3

}
