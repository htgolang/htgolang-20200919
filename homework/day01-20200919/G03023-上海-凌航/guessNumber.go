package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	// 定义次数
	rand.Seed(time.Now().Unix()) // 设置随机数种子

	rnd :=rand.Intn(100) // 生成随机数
	//println(rnd) // 打印随机数

    var isStatus bool=false

    //fmt.Println(isStatus)
    //fmt.Println(rnd)

    var maxNum=5 // 定义最大猜测次数

    //fmt.Println(maxNum)


    for i :=0; i<maxNum;i++{
    	//fmt.Println("i 的值为",i)
		var guessNum int
		fmt.Print("现在进行开始猜数字的游戏,请输入0-100 你认为正确的数字：")
		fmt.Scan(&guessNum)
    	if isStatus==false {
    		if guessNum >rnd{
    			fmt.Println("猜大了")
			}
			if guessNum ==rnd{
				fmt.Println("恭喜你猜对了")
				isStatus=true
				break
			}
			if guessNum <rnd{
				fmt.Println("猜测小了")
			}
			if i ==4{
				if guessNum !=rnd{
					fmt.Println(" 你太笨了，猜了5次都不对，程序即将退出")
					break
				}
			}
		}

	}
   }

