package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var count int = 1 	//计数器，最大输入次数为5
	var randNum int
	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(100)	//生成随机数
	fmt.Println("======猜数字游戏，数字在0-100间，您有5次机会，请输入数字======")
	fmt.Println("生成的随机数是：", randNum)
	var userNum int				//用户输入的数

	for {
		fmt.Println("请输入您猜测的数:")
		fmt.Scan(&userNum)
		if count <= 5 {
			if userNum == randNum {
				//经过count次，猜对了
				fmt.Printf("经过%d次终于猜对了，太聪明了!\n", count)
				break
			} else if userNum > randNum {
				//猜大了，还算几次机会
				count ++ 
				fmt.Printf("猜大了！还有%d次机会",5-count+1)
			} else {
				//猜小了，还算几次机会
				count ++
				fmt.Printf("猜小了！还有%d次机会",5-count+1)

			}
		} else {
			fmt.Println("5次用完还没猜对，游戏结束")
			break
		}
			
	}
}
