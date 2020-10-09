package main

import "fmt"

func main() {
	sli := []int{106, 107, 105, 109, 109, 108, 103, 102}
	//猴子掰包谷
	maxNum, secondNum := sli[0], sli[0]
	for _, v := range sli {
		if v > maxNum {
			secondNum = maxNum
			maxNum = v
		} else if v == maxNum {
			continue
		} else if v > secondNum {
			secondNum = v
		}
	}
	fmt.Println("最大数，次大数：", maxNum, secondNum)
}
