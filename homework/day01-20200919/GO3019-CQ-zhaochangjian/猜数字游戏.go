package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 随机数种子
	var num int = rand.Intn(100) // 生成0-100随机数
	fmt.Println(num)

	for i := 1; i <= 5; i++ {
		// 从标准输入接收数字
		var text int
		fmt.Print("请输入0-100之间的整数：")
		fmt.Scan(&text)

		flag := false // 退出标记

		switch {
		case text > num:
			fmt.Println("Too biger")

		case text < num:
			fmt.Println("Too smaller")

		case text == num:
			fmt.Println("Great")
			flag = true

		}
		// 猜中数字时退出
		if flag == true {
			break
		}
		// 猜了5次也退出循环
		if i == 5 {
			fmt.Println("You stupid")
			break
		}
	}

}
