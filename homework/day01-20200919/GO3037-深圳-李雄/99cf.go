package main

import "fmt"

func main() {
	// 定义初始化值，并自增
	for h := 1; h <= 9; h++ {
		// 输出换行符
		fmt.Println("")
		// 与上一个循环对比，并输出相乘的结果； 自增，并小于外层循环体的值
		for s := 1; s <= h; s++ {
			// 格式化输出 数字的形式
			fmt.Printf("%d * %d = %d\t", h, s, (h * s))
		}
	}
	fmt.Println("")
}
