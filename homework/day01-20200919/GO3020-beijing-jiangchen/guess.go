package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//Guess ...
func Guess() {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(100)
	fmt.Println("====猜数字游戏，真实数字可能是0到100之间任意的数字,共有5次机会====")
	count := 5
	var inputstr string
	for {
		fmt.Print("请输入数字:")
		_, err := fmt.Scanf("%s", &inputstr)
		if err != nil {
			fmt.Println("输入格式错误！请重新输入")
		} else {
			if input, err := strconv.Atoi(inputstr); err == nil {
				if input > number {
					count--
					if count == 0 {
						fmt.Printf("太笨了, 真实数字是%d, 程序结束\n", number)
						break
					}
					fmt.Printf("输入的数字大于真实数字, 还有%d次机会\n", count)
				} else if input < number {
					count--
					if count == 0 {
						fmt.Printf("太笨了, 真实数字是%d, 程序结束\n", number)
					}
					fmt.Printf("输入的数字小于真实数字, 还有%d次机会\n", count)
				} else {
					fmt.Println("成功")
					break
				}
			} else {
				fmt.Println("输入数字非整形, 请重新输入")
			}
		}
	}
}
