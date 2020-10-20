package main

import (
	"fmt"
	"testpkg/math" // 目录

	"mymath"
)

func main() {
	fmt.Println(math.Add(1, 2)) // 包名调用

	mymath.Test()
}
