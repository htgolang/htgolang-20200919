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

LABELOUTER:
	for {
		rand.Seed(time.Now().Unix())
		randomNum = rand.Intn(100)
	LABELINNER: //break 和 continue 后可以指定 label 用于指定跳出或跳过指定 label 同层级的循环
		for {
			for i := 1; i <= 5; i++ {
				fmt.Print("请输入猜测的值：")
				fmt.Scan(&num)
				switch {
				case num == randomNum:
					fmt.Println("1")
					break LABELINNER
				case num > randomNum:
					fmt.Println("too bigger")
				case num < randomNum:
					fmt.Println("too small")
				}
				if i == 5 {
					fmt.Printf("no chance , stop, vlue is [%d\n]", randomNum)
				}
			}
			fmt.Print("run again,[yes/y/Y]:")
			fmt.Scan(&isContinue)
			switch isContinue {
			case "yes", "y", "Y":
				fmt.Println("next")
			default:
				break LABELOUTER
			}
		}
	}
}
