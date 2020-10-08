package main

import "fmt"

func bubble(nums []int) {
	for j := 0; j < len(nums)-1; j++ {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
}

func main() {
	nums := []int{5, 4, 3, 2}
	// 最大的元素排到末尾
	// 两两比较(0->1, 1->2, 2->3) 前面大，交换
	// 两两比较，比较多少次

	// 4, 3, 2, 5
	bubble(nums)
	fmt.Println(nums)

	nums = []int{100, 80, 30, 90, 1000}
	bubble(nums)
	fmt.Println(nums)

	// b := nums
	// //b[0] = 1000
	// b[0], b[1] = b[1], b[0]
	// nums
}
