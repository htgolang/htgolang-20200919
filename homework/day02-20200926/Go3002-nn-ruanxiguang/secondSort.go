package main

import (
	"fmt"
)

// 针对问题3, 将第二大的移动到切片的最后第二位
func main() {
	var scores = []int{108, 107, 105, 109, 103, 102}
	for j := 0; j < 2; j++ {
		for i := 0; i < len(scores)-1; i++ {
			if scores[i] > scores[i+1] {
				scores[i], scores[i+1] = scores[i+1], scores[i]
			}
		}
	}
	fmt.Println(scores)
}
