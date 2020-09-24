package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number int
	rand.Seed(time.Now().Unix())
	number = rand.Int() % 100
	// fmt.Println(number)
	for i := 5; i > 0; i-- {
		var v int
		fmt.Print("请猜一下这个值的大小(0-100):")
		fmt.Scan(&v)
		if i == 1 {
			fmt.Println("<<你太笨了,游戏结束>>")
			break
		} else if v == number {
			fmt.Println("猜中了")
			break
		} else if v < number {
			fmt.Printf("太小了，还剩%d次猜测机会\n", i-1)
		} else {
			fmt.Printf("太大了，还剩%d次猜测机会\n", i-1)
		}
	}
}
