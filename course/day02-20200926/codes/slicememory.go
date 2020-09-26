package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 6)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 7)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 7)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 7)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 7)
	fmt.Println(len(nums), cap(nums))
	nums = append(nums, 7)
	fmt.Println(len(nums), cap(nums))
}
