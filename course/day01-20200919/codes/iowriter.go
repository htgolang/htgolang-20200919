package main

import (
	"fmt"
)

func main() {
	fmt.Println("1.我叫kk") // 打印内容后会自动加换行
	fmt.Println("2.我叫kk")
	fmt.Print("3.我叫kk") // 打印内容后会不加
	fmt.Print("4.我叫kk")

	name := "kk"                 // 字符串的占位是%s
	fmt.Printf("5.我叫%s\n", name) // 通过占位定义标量填充
	fmt.Printf("6.我叫%s\n", name) // 通过占位定义标量填充
}
