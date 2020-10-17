package main

import (
	"fmt"
	"testmod/math"
) // 模块名称/目录路径

func main() {
	fmt.Println("hello")
	fmt.Println(math.Add(1, 2)) // 包名+函数名称
}
