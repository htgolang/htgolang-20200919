package main

import (
	"fmt"
)

func main() {
	slice00 := []int{108, 107, 105, 109, 103, 102}

	//	// 01 bubblesort
	//	for j := 0; j < len(slice00)-1; j++ {
	//		fmt.Println("============================循环计数：", j+1)
	//		for i := 0; i < len(slice00)-1-j; i++ {
	//			fmt.Println("比较大小：", slice00[i], slice00[i+1])
	//
	//			if slice00[i] > slice00[i+1] {
	//				fmt.Println("交换位置：---> ", slice00[i], slice00[i+1])
	//				slice00[i], slice00[i+1] = slice00[i+1], slice00[i]
	//			}
	//			fmt.Println("换位结果：", slice00)
	//		}
	//		fmt.Println("循环结束，结果为:", slice00)
	//	}

	// 02 bubblesort 优化版
	for i := len(slice00) - 1; i > 0; i-- {
		fmt.Println("+++++++++++++++++++++++++++++循环计数：", i)
		for j := 0; j < i; j++ {
			fmt.Println("比大小：", slice00[j], slice00[i])
			if slice00[j] > slice00[i] {
				fmt.Println("交换位置：---> ", slice00[j], slice00[i])
				slice00[j], slice00[i] = slice00[i], slice00[j]
				fmt.Println("换位结果：", slice00)
			}
		}
		fmt.Println("循环结束，结果为:", slice00)
	}
}

//https://www.runoob.com/w3cnote/bubble-sort.html
