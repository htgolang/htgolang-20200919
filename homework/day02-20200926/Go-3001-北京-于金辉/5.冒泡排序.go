package main

import "fmt"

func bubbleSort(slice []int, order bool) {
	// 冒泡排序
	// order: true 正序, false 倒序
	flag := false
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if (slice[j] > slice[j+1] && order) || (slice[j] < slice[j+1] && !order) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				flag = true
			}
		}
		// 这是针对已经排好排的切片优化
		// flage值未变，说明没有数据进行两两交换，说明数据已经是有序状态，提前退出循环
		if !flag {
			break
		}
	}
}

func main() {
	numbers := []int{108, 107, 105, 109, 103, 102}
	//numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	bubbleSort(numbers, true)
	fmt.Println(numbers)
}
