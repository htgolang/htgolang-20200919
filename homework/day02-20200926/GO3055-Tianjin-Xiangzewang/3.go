package main

import (
	"fmt"
)

func main() {
//	var max int = 0
	nums := []int{108, 107, 105, 109, 103, 102}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
	fmt.Println(nums)
}
