package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1.猜数字

func main() {
	var index int = 0
	var number int
	rand.Seed(time.Now().Unix())
	var success_number int = rand.Intn(100)
	for {
		index += 1
		if index == 6 {
			fmt.Println("fail, try 5,exit now")
			break
		}
		fmt.Printf("请输入数字(%d):", success_number)
		fmt.Scan(&number)
		fmt.Println("\n")
		if success_number == number {
			fmt.Print("success")
			break
		} else {
			fmt.Printf("fail ... you only have %d choice", 5-index)
			fmt.Println("\n")
			continue
		}
	}

}

// 2.乘法口诀
// func main() {
// 	for i := 1; i < 10; i++ {
// 		for obj := 1; obj < i+1; obj++ {
// 			fmt.Printf("%d*%d=%d\t", i, obj, i*obj)
// 		}
// 		fmt.Println("\n")
// 	}
// }
