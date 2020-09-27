package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nums2 := []int{3, 4, 5}

	nums = append(nums, 100, 101, 102, 103)
	// for _, v := range nums2 {
	// 	nums = append(nums, v)
	// }
	// fmt.Println(nums, nums2)
	// 解包 => 切片操作
	nums = append(nums, nums2...)
	// nums = append(nums, nums2[0], nums2[1], nums2[2], ..., nums2[len()-1])
	fmt.Println(nums, nums2)

	nums = append(nums[:2], nums[3:]...)
	fmt.Println(nums)
}
