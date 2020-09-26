package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //设置随机数种子
	var randNum int = rand.Intn(100)

	var isEnd bool = false

	for i := 0; i < 5; i++ {
		var inputName int
		fmt.Print("请输入一个数字: ")
		fmt.Scan(&inputName)
		if inputName > randNum {
			fmt.Println("大了")
			isEnd = true
			continue
		} else if inputName < randNum {
			fmt.Println("小了")
			isEnd = true
			continue
		} else {
			fmt.Println("对了")
			isEnd = false
			break
		}
	}
	if isEnd {
		fmt.Println("太笨了, 5次机会已用完")
	}
}
