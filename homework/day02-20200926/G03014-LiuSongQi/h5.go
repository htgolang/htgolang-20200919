package main

import (
	"fmt"
)

/*
5. 使用冒泡排序算法对问题2切片进行从小到大排序
*/

func main() {
	var nums = []int{108, 107, 105, 109, 103, 102}

	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[x] > nums[y] {
				nums[x], nums[y] = nums[y], nums[x]
			}
		}
	}

	fmt.Println(nums)
}
