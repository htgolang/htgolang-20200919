package funcs

import "fmt"

func FindMax() {
	var nums = []int {108, 107, 105, 109, 103, 102}

	for x := 0; x < len(nums); x++ {
		for y := x +1; y < len(nums); y++ {
			if nums[x] < nums[y] {
				nums[x], nums[y] = nums[y], nums[x]
			}
		}
	}
	fmt.Println("切片中最大的值为：", nums[0])
	fmt.Println(nums)
}

func FindMaxSort() {
	var nums = []int {108, 107, 105, 109, 103, 102}

	for x := 0; x < len(nums); x++ {
		for y := x +1; y < len(nums); y++ {
			if nums[x] > nums[y] {
				nums[x], nums[y] = nums[y], nums[x]
			}
		}
	}
	fmt.Println("切片中最大的值为：", nums[len(nums)-1])
	fmt.Println(nums)
}

func FindMaxFinal() {
	var nums = []int {108, 107, 105, 109, 103, 102}

	for x := 0; x < len(nums); x++ {
		for y := x +1; y < len(nums); y++ {
			if nums[x] < nums[y] {
				nums[x], nums[y] = nums[y], nums[x]
			}
		}
	}
	fmt.Println("切片中最大的值为：", nums[0])
	fmt.Println(nums)
	nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
	fmt.Println(nums)
}