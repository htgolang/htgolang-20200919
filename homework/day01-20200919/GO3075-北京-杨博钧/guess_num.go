package main

import (
	"fmt"
	"time"
	"math/rand"
)

func GuessNum() {
	//定义label使猜错5次后重新开始
	START:
	//生成随机种子
	rand.Seed(time.Now().UnixNano())
	//定义变量分别是这次生成的随机数、猜测次数、用户猜测的值
	num := rand.Intn(100)
	count := 0
	guess := -1
	fmt.Printf("猜数字游戏，请输入一个0~99的数字看看是否正确:\n")
	//开启for循环开始猜测
	LABEL:
	for {
		//接收猜测的数字
		fmt.Scanln(&guess)
		count++
		//判断接收到的数字是否符合范围
		if guess < 0 || guess >= 100 {
			fmt.Printf("当前输入的%v不满足0~99整数范围，请重新输入一个:\n", guess)
			continue
		}
		//进行判断
		if guess == num && count < 5 {
			fmt.Println("猜对了！")
			break
		} else if guess < num && count < 5 {
			fmt.Printf("猜小了，再猜一个试试吧，当前还有%v次机会:\n", 5 - count)
		} else if guess > num && count < 5 {
			fmt.Printf("猜大了，再猜一个试试吧，当前还有%v次机会:\n", 5 - count)
		} else {
			fmt.Printf("太笨了！\n")
			//选择是否重新开始
			fmt.Printf("是否重新开始游戏?(y/n):\n")
			var choice string
			fmt.Scanln(&choice)
			switch choice {
			case "y", "Y" :
				goto START
			case "n", "N" :
				break LABEL
			default:
				fmt.Println("请输入正确的选项")
			}
		}
	}
}

func main() {
	GuessNum()
}