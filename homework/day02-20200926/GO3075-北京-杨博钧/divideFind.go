package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	slice = []int{108, 107, 105, 109, 103, 102}
)

func InsertSort(this []int) {
	for j := 1; j < len(this); j++ {
		for i := j; i > 0 && this[i] < this[i - 1]; i-- {
			this[i], this[i - 1] = this[i - 1], this[i]
		}
	}
}


//习题七，对排序后的切片二分查找
func divideSearch(data int, arr []int) int {
	if len(arr) == 1 && arr[0] != data {
		return -1
	}
	if data == arr[len(arr) / 2] {
		return len(arr) / 2
	} else if data < arr[len(arr) / 2] {
		if divideSearch(data, arr[:len(arr)/2]) > 0 {
			return divideSearch(data, arr[:len(arr)/2])
		} else {
			return divideSearch(data, arr[:len(arr)/2])
		}
	} else if data > arr[len(arr) / 2] {
		if divideSearch(data, arr[len(arr)/2:]) > 0 {
			return len(arr)/2 + divideSearch(data, arr[len(arr)/2:])
		} else {
			return divideSearch(data, arr[len(arr)/2:])
		}
	}
	return 0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numRand := rand.Intn(10) + 100
	//使用插入排序使切片有序
	InsertSort(slice)
	//打印随机数和排序后切片
	fmt.Printf("本次生成的随机数为：%v\n", numRand)
	fmt.Printf("切片为：%v\n", slice)
	//二分查询
	p := divideSearch(numRand, slice)
	if p < 0 {
		fmt.Printf("没找到，切片内没有%v\n", numRand)
	} else {
		fmt.Printf("找到%v了，在切片内下标为%v\n", numRand, p)
	}
}