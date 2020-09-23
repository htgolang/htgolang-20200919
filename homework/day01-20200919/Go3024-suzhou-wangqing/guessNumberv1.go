/*
需求
猜数字 生成随机数字0-100 从控制台数据 与生成数字比较 大 提示太大了 小 提示太小了 等于 成功, 程序结束
最多猜测五次，未猜对，说太笨了，程序结束
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number int
	number = rand.Intn(100)
	// fmt.Println(number)
	var gnumber int

	for times := 1; times <= 5; times++ {
		fmt.Print("请输入一个数字:")
		fmt.Scan(&gnumber)
		fmt.Println("你输入的数字是：", gnumber)

		switch {
		case gnumber > number:
			fmt.Println("太大了")
		case gnumber < number:
			fmt.Println("太小了")
		default:
			fmt.Println("成功了")
		}

		if times == 5 {
			fmt.Println("你已经猜了5次了，游戏结束")
		}

	}

	// 	fmt.Println("你太笨了")

}
