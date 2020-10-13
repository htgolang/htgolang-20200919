package main

import "fmt"

func main() {
	var slice1 = []int{108, 135, 166, 107, 352, 666, 105, 242, 112, 109, 103, 102}

	// 1
	// fmt.Println(maxValue(slice1))

	// 2
	// fmt.Println(maxEnd(slice1))

	// 3
	// fmt.Println(secondEnd(slice1))

	// 4
	fmt.Println(bubbleSort(slice1))
}

// 	1.取切片中最大值
func maxValue(slice []int) (int, int) {
	maxIndex, maxValue := 0, 0

	for i := 0; i < len(slice); i++ {
		if slice[i] > maxValue {
			maxIndex = i
			maxValue = slice[i]
		}
	}

	return maxIndex, maxValue
}

// 2.将最大值移动至最后一位
func maxEnd(slice []int) []int {
	maxIndex, maxValue := maxValue(slice)

	slice = append(slice[:maxIndex], slice[maxIndex+1:]...)

	slice = append(slice, maxValue)

	return slice
}

// 3.将第二大移动至倒数第二位
func secondEnd(slice []int) []int {
	_, maxValue := maxValue(slice)

	slice2 := maxEnd(slice)

	slice2 = maxEnd(slice2[:len(slice2)-1])

	slice2 = append(slice2, maxValue)

	return slice2
}

// 4.泡沫排序
func bubbleSort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j+1] < slice[j] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}
