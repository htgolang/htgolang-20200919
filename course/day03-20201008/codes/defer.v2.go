package main

import "fmt"

func main() {
	fmt.Println("start")

	// defer 函数调用
	// defer 延迟执行
	// 在函数退出之前执行
	for i := 0; i < 2; i++ {
		defer func(i int) {
			fmt.Printf("defer %d \n", i)
		}(i)
	}

	// defer 0 defer 1
	fmt.Println("end")
	// start -> end -> defer 1 -> defer 0 => 堆栈
}
