package main

import (
	"fmt"
)

// int切片 []int{108, 107, 105, 109, 103, 102},找出最大的数字
func main() {
	var scores = []int{108, 107, 105, 109, 103, 102}
	var max = scores[0]
	for i := 0; i < len(scores); i++ {
		if max < scores[i] {
			max = scores[i]
		}
	}
	fmt.Println(max)
}
