package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close() //?

	// fmt.Println(file.WriteString("123"))
	writer := bufio.NewWriter(file)
	fmt.Println(writer.WriteString("123123123"))
	fmt.Println(writer.Flush())
}
