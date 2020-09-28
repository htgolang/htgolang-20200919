package main

import (
	"fmt"
	"math/rand"
)

var shuzi int
var suiji int = rand.Intn(100)

func main() {
	fmt.Println(suiji)
	for i := 1; i <= 5; i++ {
		fmt.Printf("请输入一个数字:")
		fmt.Scan(&shuzi)
		//等待用户输入一个数字
		if i > 4 {
			fmt.Println("猜了5次还没猜对，你太笨了！！")
			break
		}
		if shuzi > suiji {
			fmt.Println("数字太大了，重新来过")
		} else if shuzi < suiji {
			fmt.Println("数字太小了，重新来过")
		} else if shuzi == suiji {
			fmt.Println("太好了，猜对了")
			break
		}
	}
	fmt.Println("随机数字为:", suiji)
}
