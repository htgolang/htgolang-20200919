package main

import "fmt"

func main() {
	nums := []int{108, 107, 105, 109, 103, 102}
	//将最大的数字移动到切片的最后一位 原来的数字都在移动后的切片中都存在,
	//只是最大的数字再最后一位
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
	fmt.Println(nums)
	//将第二大的移动到切片的最后第二位
	nums = []int{108, 107, 105, 109, 103, 102}
	for i := len(nums) - 2; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
	//冒泡排序
	nums = []int{108, 107, 105, 109, 103, 102}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println(nums)
}
