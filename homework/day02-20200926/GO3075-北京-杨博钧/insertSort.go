package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	slice = []int{108, 107, 105, 109, 103, 102}
)

//习题六、插入排序
//每一步将一个待排序的数据插入到前面已经排好序的有序序列中，直到插完所有元素为止
func InsertSort(this []int) {
	for j := 1; j < len(this); j++ {
		for i := j; i > 0 && this[i] < this[i - 1]; i-- {
			this[i], this[i - 1] = this[i - 1], this[i]
		}
	}
}

func main() {
	InsertSort(slice)
	fmt.Println(slice)
}