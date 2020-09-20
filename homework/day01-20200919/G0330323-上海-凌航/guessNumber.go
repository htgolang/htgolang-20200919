package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 定义次数
	rand.Seed(time.Now().Unix()) // 设置随机数种子

	rnd := rand.Intn(100) // 生成随机数

	//println("====",rnd) // 打印正确的数字

	for count := 1; count < 7; count++ { // 初始值为1
		if count <= 5 {
			var guessNum int
			fmt.Print("现在进行开始猜数字的游戏,请输入0-100 你认为正确的数字：")
			fmt.Scan(&guessNum)
			if guessNum > rnd {
				fmt.Println("数字猜大了")
			} else if guessNum == rnd {
				fmt.Println("恭喜你猜对了")
				break
			} else if guessNum < rnd {
				fmt.Println(" 数字猜小了")
			}
		}
		if count > 5 { // 如果大于5次程序结束，打印提示
			fmt.Println("最多猜测五次，未猜对，太笨了，程序结束")
			break
		}
	}
}
