package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	var choice string
	fmt.Println("1.table", "2.guess")
	fmt.Scan(&choice)	
	switch choice {
	case "1":
		Table()
	//default:
	case "2":
		Guess()
	}
}

func Table(){
	fmt.Println("======打印九九乘法表======")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			//fmt.Printf("%v * %v = %2v\t", i, j, i*j)
			fmt.Printf("%d * %d = %2d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println("=======打印结束======")
}

func Guess() {
	var count int = 1 //计数器，最大输入次数为5
	var randNum int
	rand.Seed(time.Now().Unix())
	randNum = rand.Intn(100) //生成随机数
	fmt.Println("======猜数字游戏，数字在0-100间，您有5次机会，请输入数字======")
	fmt.Println("生成的随机数是：", randNum)
	var userNum int //用户输入的数,默认为字符串，需类型转换

	for {
		fmt.Println("请输入您猜测的数:")
		_, err := fmt.Scan(&userNum)
		fmt.Println(err)
		if err != nil {
			fmt.Println("格式错误，请重新输入！")
			break
		} else {
			if count <= 5 {
				if userNum == randNum {
					//经过count次，猜对了
					fmt.Printf("经过%d次终于猜对了，太聪明了!\n", count)
					break
				} else if userNum > randNum {
					//猜大了，还算几次机会
					count++
					fmt.Printf("猜大了！还有%d次机会", 5-count+1)
				} else {
					//猜小了，还算几次机会
					count++
					fmt.Printf("猜小了！还有%d次机会", 5-count+1)

				}
			} else {
				fmt.Println("5次用完还没猜对，游戏结束")
				break
			}
		}
	}
}