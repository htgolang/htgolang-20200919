package main

import (
	"fmt"
)

var (
	slice = []int{108, 107, 105, 109, 103, 102}
)
//习题二，从切片中获取最大值
func GetSliceMax(this []int) (max int) {
	for _, v := range this {
		if v > max {
			max = v
		}
	}
	return
}

func main() {
	max := GetSliceMax(slice)
	fmt.Println(max)
}