package main

//int切片 []int{108, 107, 105, 109, 103, 102} 找出最大的数字

import (
	"fmt"
)

func main() {
	num_list := []int{108, 107, 105, 109, 103, 102}
	var maxNum int

	for i := 0; i < len(num_list)-1; i++ {
		if maxNum < num_list[i] {
			maxNum = num_list[i]
		}

	}
	fmt.Println(maxNum)

}
