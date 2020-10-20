package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		3.针对问题2的切片, 将最大的数字移动到切片的最后一位 原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位
		4.针对问题3, 将第二大的移动到切片的最后第二位

	*/

	// 定义int空切片
	var nu_list []int

	// 初始化切片
	nu_list = []int{108, 107, 105, 109, 103, 102, 110, 106}
	// fmt.Println(nu_list[7])
	// nu_list = []int{108, 107, 105, 109, 103, 102, 106, 130}
	for j := 0; j < 2; j++ {
		for i := 0; i < len(nu_list)-1; i++ {
			if nu_list[i] > nu_list[i+1] {
				nu_list[i], nu_list[i+1] = nu_list[i+1], nu_list[i]
			}

			fmt.Printf("第 %d 轮:\t数据：%d\n", i, nu_list)
		}
		fmt.Println(strings.Repeat("-", 60))
	}

}
