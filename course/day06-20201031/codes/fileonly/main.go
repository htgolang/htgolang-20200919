package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content := make([]byte, 10)
	fmt.Println(file.Read(content))
	fmt.Println(string(content))
	fmt.Println(file.Write([]byte("123")))

}
