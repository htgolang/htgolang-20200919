package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randomNum int
	var num int
	var isContinue string

	rand.Seed(time.Now().Unix())
	fmt.Println()
	randomNum = rand.Intn(100)
	fmt.Println(randomNum)
//OuterLoop:
	for {
		for i := 1; i <= 5; i++ {
			fmt.Print("请输入猜测的值：")
			fmt.Scan(&num)
			if num == randomNum {
				println("猜对了")
				break
			} else if num > randomNum {
				if num > 100 {
					println("请输入0~100之间的数字")
					i--
					fmt.Println(i)
					continue
				}
				fmt.Println("数字太大了")
			} else if num < randomNum {
				if num <= 0 {
					fmt.Println("请输入0~100之间的数字")
					i--
					fmt.Println(i)
					continue
				}
				fmt.Println("数字太小了")
			}
			fmt.Println(i)
			//switch {
			//case num == randomNum:
			//	fmt.Println("1")
			//	break LABELINNER
			//case num > randomNum:
			//	fmt.Println("too bigger")
			//case num < randomNum:
			//	fmt.Println("too small")
			//}
			//if i == 5 {
			//	fmt.Printf("no chance , stop, vlue is [%d\n]", randomNum)
			//}
		}
		fmt.Print("run again,[yes/y/Y]:")
		fmt.Scan(&isContinue)
		switch isContinue {
		case "yes", "y", "Y":
			fmt.Println("next")
		default:
			//break OuterLoop
			return
		}
	}
}
