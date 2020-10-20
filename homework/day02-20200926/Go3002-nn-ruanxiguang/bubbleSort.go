package main

import (
	"fmt"
)

// 使用冒泡排序算法对问题2切片进行从小到大排序
func main() {
	var scores = []int{108, 107, 105, 109, 103, 102}
	for j := 0; j < len(scores); j++ {
		for i := 0; i < len(scores)-1; i++ {
			if scores[i] > scores[i+1] {
				scores[i], scores[i+1] = scores[i+1], scores[i]
			}
		}
	}
	fmt.Println(scores)
}