package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var slice1 = []int{108, 107, 105, 109, 103, 102}

	var slice2 []int

	// 将生成的随机数插入到切片中
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(110)
		slice2 = append(slice2, randNum)
	}

	// 判断元素是否在切片中
	for _, v2 := range slice2 {
		iscontainers, i := isContain(insertSort(slice1), v2)
		if iscontainers {
			fmt.Printf("包含%d,索引是%d\n", v2, i)
		} else {
			fmt.Printf("不包含%d\n", v2)
		}
	}

	// fmt.Println(insertSort(slice1))
}

func insertSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func isContain(slice []int, item int) (bool, int) {
	for i, v := range slice {
		if v == item {
			return true, i
		}
	}
	return false, 0
}
