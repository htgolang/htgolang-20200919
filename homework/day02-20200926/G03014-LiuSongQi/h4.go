package main

import "fmt"

/*
针对问题3, 将第二大的移动到切片的最后第二位
*/

func main() {
	var nums = []int{108, 107, 105, 109, 103, 102}
	var maxIndex int
	var subMaxIndex int
	for x, y := range nums {
		if nums[maxIndex] < y {
			subMaxIndex = maxIndex
			maxIndex = x
		} else {
			if nums[subMaxIndex] < y {
				subMaxIndex = x
			}
		}
	}

	var sortNums = []int{}
	for i, v := range nums {
		if i != maxIndex && i != subMaxIndex {
			sortNums = append(sortNums, v)
		}
	}

	sortNums = append(sortNums, nums[subMaxIndex], nums[maxIndex])
	fmt.Println(sortNums)
}
