package main

import "fmt"

func main() {
	nums := [...]int{1, 2, 3, 4, 5}
	// nums2 := nums[1:3]
	// fmt.Printf("%T\n", nums2)
	// fmt.Println(nums2)
	// //0 <= start <= end <= length
	// // cap = length - start
	// // len = end - start
	// fmt.Println(cap(nums2), len(nums2))

	// start end cap_end
	nums2 := nums[1:3:4]
	fmt.Println("%T\n", nums2)
	fmt.Println(nums2)
	// 0 <= start <= end <= cap_end <= length
	// cap = cap_end -start
	// len = end - start
	fmt.Println(cap(nums2), len(nums2))
}
