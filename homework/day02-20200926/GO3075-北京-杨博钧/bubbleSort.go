package main

import (
	"fmt"
)

var (
	slice = []int{108, 107, 105, 109, 103, 102}
)

//习题五、冒泡排序
func BubbleSort(this []int) {
	for j := 1; j < len(this); j++ {
		for i := 0; i < len(this) - j; i++ {
			if this[i] > this[i + 1] {
				this[i + 1], this[i] = this[i], this[i + 1]
			}
		}
	}
}

func main() {
	BubbleSort(slice)
	fmt.Println(slice)
}