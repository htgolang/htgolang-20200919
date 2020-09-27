package main

import "fmt"

func main() {
	var nums [2][]int
	fmt.Println(nums)
	nums[0] = append(nums[0], 1)
	nums[1] = append(nums[1], 100)
	fmt.Println(nums)

	var nums2 [][2]int
	fmt.Println(nums2)
	nums2 = append(nums2, [2]int{1, 2})
	nums2 = append(nums2, [2]int{10, 20})
	fmt.Println(nums2)

	var nums3 [][]int
	nums3 = append(nums3, []int{})
	nums3 = append(nums3, []int{})
	nums3 = append(nums3, []int{})

	nums3[0] = append(nums3[0], 1, 2, 3)
	nums3[1] = append(nums3[1], 2, 3, 4)
	fmt.Println(nums3)

	var nums4 [1][]int
	nums4[0] = []int{1, 2, 3}

	fmt.Println(nums4)

	var nums5 [0]int
	fmt.Println(nums5)
}
