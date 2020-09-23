package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randNum int				//随机数
	var userNum int				//输入数
	var isContinue string		//是否继续
	
	OuterLoop:
		for {
			rand.Seed(time.Now().Unix())
			randNum = rand.Intn(100)	//生成随机数
			fmt.Println("======猜数字游戏，数字在0-100间，您有5次机会，请输入数字======")
			fmt.Println("生成的随机数是：", randNum)
		InnerLoop:
			for i := 1; i <= 5; i++ {
				fmt.Println("请输入您猜测的数:")
				fmt.Scan(&userNum)
				switch {
				case userNum == randNum:
				
					fmt.Printf("经过%d次猜对了，太聪明了!\n", i)
					break InnerLoop
				case userNum > randNum:
					fmt.Println("你猜大了")
				case userNum < randNum:
					fmt.Println("你猜小了")
				}
				if i == 5 {
					fmt.Printf("5次机会用完了")
				}
			}
			fmt.Println("继续猜，[yes/y/Y]")
			fmt.Scan(&isContinue)
			switch isContinue {
			case "yes","y","Y":
				fmt.Println("继续")
			default:
				break OuterLoop
			}	
		}
	
}

