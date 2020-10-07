package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums2 := nums[1:3]
	fmt.Println(cap(nums2))
	nums2 = append(nums2, 100)
	fmt.Println(nums2)
	fmt.Println(nums)
}
