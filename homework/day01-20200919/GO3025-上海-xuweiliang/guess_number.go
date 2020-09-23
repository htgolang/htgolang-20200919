package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_GUESS = 5

func main() {
	var answer int

	rand.Seed(time.Now().Unix())
	result := rand.Int() % 100

	for i := 1; i <= MAX_GUESS; i++ {
		fmt.Print("一共有5次机会，请输入你猜的数字：")

		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Println("请重新输入")
			continue
		}
		if answer > result {
			fmt.Printf("猜大了，还有%d次机会\n", MAX_GUESS-i)
		} else if answer < result {
			fmt.Printf("猜小了，还有%d次机会\n", MAX_GUESS-i)
		} else {
			fmt.Println("猜对了，太棒了！")
			break
		}
	}
}
