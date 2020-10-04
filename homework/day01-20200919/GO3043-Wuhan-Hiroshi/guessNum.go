package main

import (
	"fmt"
	"math/rand"
	"time"
)

var num int

func main() {
	rand.Seed(time.Now().Unix())
	randomint := (rand.Intn(100))
	fmt.Println(randomint)

	// 做每次输入的次数累加，初始值为1，当输入5次都不正确后，i累加变成6,此时循环也结束;
	// 但是正确时不能做累加,如果刚好第5次结果正确，累加i的值会变成6
	i := 1

	for n := 1; n <= 5; n++ {
		fmt.Print("Input like num: ")
		fmt.Scan(&num)
		if num > randomint {
			fmt.Println("输入的数字大了点,剩余机会:", 5-n)
			i++
		} else if num < randomint {
			fmt.Println("输入的数字小了点,剩余机会:", 5-n)
			i++
		} else {
			fmt.Println("刚好啦!")
			break
		}
	}
	if i > 5 {
		fmt.Println("机会用完啦,运气不佳哦!")
	}
}
