package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number int
	var status bool = false

	// 生成一个随机数字
	rand.Seed(time.Now().Unix())
	randint := rand.Intn(100)
	fmt.Println(randint)

	for i := 0; i <= 5; i++ {
		fmt.Print("猜猜我的数字是啥? >>> ")
		fmt.Scan(&number)
		if number > randint {
			fmt.Println("太大了")
		} else if number < randint {
			fmt.Println("太小了")
		} else {
			fmt.Println("恭喜你，猜对了!")
			status = true
			break
		}
	}

	if !status {
		fmt.Println("太笨了")
	}
}
