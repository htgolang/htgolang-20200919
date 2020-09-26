package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums[0:3])
	fmt.Println(nums[:3])
	fmt.Println(nums[3:len(nums)])
	fmt.Println(nums[3:])
	fmt.Println(nums[:])
}
