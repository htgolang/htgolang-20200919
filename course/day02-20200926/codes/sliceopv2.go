package main

import "fmt"

func main() {
	nums := make([]int, 5, 10) // len=5, cap=100
	// 0 0 0 0 0[4] 5
	// slice[start:end]
	// 0 <= start <= end <= cap
	// len := end-start
	// cap := cap-start
	// nums2 := nums[2:5]

	// fmt.Println(nums2, len(nums2), cap(nums2))

	// slice[start:end:cap_end]
	// 0 <= start <= end <= cap_end <= cap
	// len := end - start
	// cap := cap_end - start
	nums2 := nums[2:5:5]
	fmt.Println(nums2, len(nums2), cap(nums2))
	nums2 = append(nums2, 1000)
	nums = append(nums, 1)
	fmt.Println(nums2, nums)
}
