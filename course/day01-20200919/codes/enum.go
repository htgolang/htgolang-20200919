package main

import "fmt"

func main() {
	// 枚举值
	// iota 在一个小括号内, 初始化为0, 每调用一次+1
	// const (
	// 	statusNew      = iota // 0
	// 	statusComplete = iota // 1
	// 	statusDeleted  = iota // 2
	// )

	const (
		statusNew      = iota // 0
		statusComplete        // 1
		statusDeleted         // 2
	)

	const (
		Monday = iota * 10
		Tuesday
		Wednesday
	)

	fmt.Println(statusNew, statusComplete, statusDeleted)

	fmt.Println(Monday, Tuesday, Wednesday)
}
