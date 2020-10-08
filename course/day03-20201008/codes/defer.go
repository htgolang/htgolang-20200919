package main

import "fmt"

func main() {
	fmt.Println("start")

	// defer 函数调用
	// defer 延迟执行
	// 在函数退出之前执行
	defer func() {
		fmt.Println("defer")
	}()

	fmt.Println("end")
	// start -> end -> defer
}
