package main

import "fmt"

func main() {
	nums := [...]int{108, 107, 105, 109, 103, 102}

	for j := 0; j < len(nums)-1; j++ {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	fmt.Println(nums)
}
