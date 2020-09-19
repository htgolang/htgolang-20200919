package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入你的名字：")
	fmt.Scan(&name) // 接收控制台输入内容 赋值给变量name
	// &name => 取name指针(地址)
	fmt.Println("你输入的内容是:", name)

	var age int

	fmt.Print("请输入你的年龄:")
	fmt.Scan(&age)
	fmt.Println("你输入的年龄是:", age)
}
