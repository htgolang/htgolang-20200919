package main

import (
	"fmt"
)

// 针对问题2的切片, 将最大的数字移动到切片的最后一位
// 原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位
func main() {
	var scores = []int{108, 107, 105, 109, 103, 102}
	for i := 0; i < len(scores)-1; i++ {
		if scores[i] > scores[i+1] {
			scores[i], scores[i+1] = scores[i+1], scores[i]
		}
	}
	fmt.Println(scores)
}
