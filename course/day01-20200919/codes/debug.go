package main

import "fmt"

func main() {
	var a = 1

	// 100行代码 50打印(a) 75
	a = a + 2
	a = a + 3

	fmt.Println(a) // 5
}
