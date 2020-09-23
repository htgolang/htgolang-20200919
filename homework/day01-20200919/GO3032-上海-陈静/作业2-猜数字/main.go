package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		需求：
		猜数字 生成随机数字0-100
		从控制台数据 与生成数字比较
		大 提示太大了
		小 提示太小了
		等于 成功, 程序结束
		最多猜测五次，未猜对，说太笨了，程序结束
	*/
	var randNum int
	rand.Seed(time.Now().Unix()) //以时间创建随机数种子
	randNum = rand.Intn(100)
	fmt.Println(randNum)

	nums := 1
	for ; nums <= 5; nums++ {
		var guessnum int
		fmt.Printf("请输入数字guesnum：")
		fmt.Scan(&guessnum)

		if guessnum > randNum {
			fmt.Println("太大了")
		} else if guessnum < randNum {
			fmt.Println("太小了")
		} else {
			fmt.Println("猜对了")
			break
		}

	}

	if nums == 6 {
		fmt.Println("未猜对，太笨了")
	}

}
