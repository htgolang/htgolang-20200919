package main

import "fmt"

func main() {
	var maxNum int = 0
	var maxIndex int
	nums := [...]int{108, 107, 105, 109, 103, 102}

	for i, v := range nums {
		if v >= maxNum {
			maxNum = v
			maxIndex = i
		}
	}
	copy(nums[maxIndex:], nums[maxIndex+1:])
	nums[len(nums)-1] = maxNum
	// fmt.Println(nums)

	maxNum = 0
	for i, v := range nums[0 : len(nums)-1] {
		if v >= maxNum {
			maxNum = v
			maxIndex = i
		}
	}

	copy(nums[maxIndex:], nums[maxIndex+1:])
	nums[len(nums)-2] = maxNum
	// fmt.Println(maxNum)
	// fmt.Println(maxIndex)
	fmt.Println(nums)
}
