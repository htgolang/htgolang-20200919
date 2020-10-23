package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
   随机数0-100
   从控制台输入数字与随机数相比大小
   1. 大了提示大了
    2. 小了提示小了
    3. 等于  成功
    4. 5次机会  没有成功输出太笨了

 */

func main() {


    for b := 1; b <=5 ; b++ {
		var b int
		fmt.Println("请输入你的数字")
		fmt.Scan(&b)
		rand.Seed(time.Now().Unix())
		a := rand.Intn(100)
		fmt.Println(a)
		if b > a {
			fmt.Println("太大了")
		} else if b < a {
			fmt.Println("太小了")
		} else if b == a {
			fmt.Println("猜对了")
			break
		}

	}
	fmt.Println("太笨了")

}
