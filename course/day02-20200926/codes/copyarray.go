package main

import "fmt"

func main() {
	// 复制nums中的所有数据到nums2(两个不会有相互影响)
	// v1
	// nums := []int{1, 2, 3, 4, 5}
	// nums2 := make([]int, len(nums))
	// copy(nums2, nums)
	// v2
	// nums := []int{1, 2, 3, 4, 5}
	// nums2 := make([]int, len(nums)) // []int{0, 0, 0, 0, 0}
	// for i, v := range nums {
	// 	nums2[i] = v
	// }
	// fmt.Println(nums2)
	nums := []int{1, 2, 3, 4, 5}
	nums2 := make([]int, 0, len(nums))
	// 空白标识符
	for _, v := range nums {
		nums2 = append(nums2, v)
	}
	fmt.Println(nums2)
}
