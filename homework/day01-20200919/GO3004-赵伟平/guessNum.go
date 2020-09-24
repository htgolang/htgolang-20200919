package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		var Num int
		for {
			fmt.Println("请输入你要猜测的数字，100 以内：")
			fmt.Scan(&Num)
			// fmt.Println(Num)
			if Num > 0 && Num < 100 {
				break
			} else {
				fmt.Println("警告，请遵守游戏规则，输入 100 以内的数字!!!")
			}
		}
		// fmt.Println(Num)
		rand.Seed(time.Now().Unix())
		randNum := rand.Intn(100)
		// fmt.Println(reflect.TypeOf(Num))
		// fmt.Println(reflect.TypeOf(randNum))
		fmt.Println(randNum)

		if Num == randNum {
			fmt.Println("恭喜你，猜对了！！！")
			break
		} else if Num < randNum {
			fmt.Println("太小了")
		} else if Num > randNum {
			fmt.Println("太大了")
		}
	}

}
