package main

import (
	"fmt"
	"math/rand"
)

func guessNum() {
	n := rand.Intn(101) // [0,101) 即0-100的随机数
	var typeIn int
	correct := false
	for i := 1; i <= 5; i++ {
		fmt.Printf("猜数字，当前第%d次\n", i)
		fmt.Scan(&typeIn)
		switch {
		case typeIn > n:
			fmt.Println("猜大了")
		case typeIn < n:
			fmt.Println("猜小了")
		default:
			correct = true
		}
		if correct == true {
			fmt.Println("猜对了，游戏结束")
			return
		}
	}
	fmt.Println("5次未猜对，游戏结束")
}

func main() {
	guessNum()
}
