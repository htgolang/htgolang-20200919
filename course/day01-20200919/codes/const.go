package main

import "fmt"

// 定义常量(需要初始化值)
const (
	statusNew     int = 1
	statusDeleted     = 2
)

func main() {
	const (
		Monday = 10
		/* 在一个小括号内
		若为赋值，则使用最近的一个已赋值的常量对应的值进行初始化
		*/
		Tuesday = 20
		Wednesday
	)

	fmt.Println(statusNew, statusDeleted)
	fmt.Println(Monday, Tuesday, Wednesday)
}
