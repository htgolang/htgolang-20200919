package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums2 := nums[1:3] // [2, 3]

	nums = append(nums, 1000) // [1, 2, 3, 4, 5, 1000] x

	nums2 = append(nums2, 100) // [2, 3, 100]
	fmt.Println(nums2)         //  [2, 3, 100]
	fmt.Println(nums)          // [1, 2, 3, 4, 5, 1000]

}
