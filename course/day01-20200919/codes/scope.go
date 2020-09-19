package main

import (
	"fmt"
)

var name string = "kk"

func main() {
	/*
		块注释
	*/
	fmt.Println(name) // 打印的kk python => ??
	// 作用域 说明标识符的使用范围 {}
	var name string = "silence" // 定义变量name
	var age int = 32

	{
		var age int = 33
		fmt.Println(name, age) //silence, 33
		name = "kk"
	}

	fmt.Println(name, age) // silence, 32
}
