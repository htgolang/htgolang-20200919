package main

import (
	"fmt"
)

func Bobble_Sort_order(nums []int) []int {

	for j := 0; j < len(nums); j++ {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}

func Bobble_Sort_positive(nums []int) []int {
	for j := 0; j < len(nums); j++ {
		for i := 1; i < len(nums)-j; i++ {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
	return nums
}

func main() {

	var list []int

	list = []int{108, 107, 105, 109, 103, 102, 110, 106}

	fmt.Println(Bobble_Sort_order(list))
	fmt.Println(Bobble_Sort_positive(list))
}
