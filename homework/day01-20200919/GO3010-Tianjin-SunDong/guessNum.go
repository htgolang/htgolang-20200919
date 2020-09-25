package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	randint := rand.Intn(100)
	fmt.Println("猜数字游戏开始")
	fmt.Println("游戏规则：系统生成的数字为100以内的整数，您有5次机会。加油！！！")

	for t := 0; t < 5; t++ {
		var testNum int
		fmt.Printf("第%v次，请输入一个100以内的数字：", t+1)
		fmt.Scan(&testNum)
		//fmt.Println(randint, testNum)

		switch {
		case testNum > randint:
			fmt.Println("猜大了！")
		case testNum < randint:
			fmt.Println("猜小了！")
		case testNum == randint:
			fmt.Println("恭喜，猜对了！")
			return
		default:
			fmt.Println("调皮！")
		}

	}
	fmt.Println("游戏结束，您没有猜到哦！！")
}
