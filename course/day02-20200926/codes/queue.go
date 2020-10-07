package main

import "fmt"

func main() {
	// 队列
	// 先进先出
	queue := []int{}

	queue = append(queue, 1)
	queue = append(queue, 3)
	queue = append(queue, 2)
	queue = append(queue, 6)

	// append 右边进入
	// 1 -> [1]
	// 2 -> [1, 2]
	// 3 -> [1, 2, 3]
	// 从左边出
	// <- 1  [2, 3]
	// <- 2 [3]
	// <- 3 []
	for len(queue) != 0 {
		fmt.Println(queue[0])
		queue = queue[1:]
	}

	fmt.Println("over")
}
