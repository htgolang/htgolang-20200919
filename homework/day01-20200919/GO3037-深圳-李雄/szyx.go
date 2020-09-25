package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_GUESS int = 5 // 定义常量

func main() {
	var answer int

	rand.Seed(time.Now().Unix()) // 使用当前时间段设置随机数种子
	result := rand.Int() % 100   // 生成【0， 100】的随机数

	for i := 1; i <= MAX_GUESS; i++ {
		fmt.Printf("系统数字：%d", result)
		fmt.Println("请输入数字：")

		if _, err := fmt.Scanln(&answer); err != nil {
			fmt.Println("请重新输入")
			continue
		}

		if answer > result {
			fmt.Printf("猜大了，还有%d机会\n", MAX_GUESS-i)
		} else if answer < result {
			fmt.Printf("猜小了，还有%d次机会\n", MAX_GUESS-i)
		} else {
			fmt.Println("猜对了， 太棒了")
			break
		}
	}
}
