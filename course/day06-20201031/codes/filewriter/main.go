package main

import (
	"fmt"
	"os"
)

func main() {
	path := "test.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println(file.Write([]byte("123456789")))
	fmt.Println(file.Write([]byte("我是kk")))
	file.WriteString("啦啦啦啦")
}
