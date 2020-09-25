package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(100)

	for c := 10; c > 0; c-- {
		var inputstr string
		fmt.Printf("请输入数字:")
		fmt.Scan(&inputstr)
		input, err := strconv.Atoi(inputstr)
		if err == nil {

			if input > number {
				fmt.Println("大了")

			} else if input < number {
				fmt.Println("小了")

			} else {
				fmt.Println("猜对了")
				break
			}
		} else {
			fmt.Println("请输入整数")
		}

	}

}
