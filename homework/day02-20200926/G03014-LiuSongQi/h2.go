package main

import (
	"fmt"
)

/*
int切片 []int{108, 107, 105, 109, 103, 102} 找出最大的数字
*/

func main() {
	var nums = []int{108, 107, 105, 109, 103, 102}

	// for x := 0; x < len(nums); x++ {
	// 	for y := x + 1; y < len(nums); y++ {
	// 		if nums[x] < nums[y] {
	// 			nums[x], nums[y] = nums[y], nums[x]
	// 		}
	// 	}
	// }

	for x := 0; x < len(nums)-1; x++ {
		if nums[x] > nums[x+1] {
			nums[x], nums[x+1] = nums[x+1], nums[x]
		}
	}

	fmt.Println(nums[len(nums)-1])
}
