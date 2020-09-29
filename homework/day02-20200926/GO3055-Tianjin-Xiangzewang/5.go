package main

import (
	"fmt"
)

func main() {
	nums := []int{108, 107, 105, 109, 103, 102}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j< len(nums)-1-i; j++ {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
	}
	fmt.Println(nums)
}
