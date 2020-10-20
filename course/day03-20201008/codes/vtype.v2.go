package main

import "fmt"

func main() {
	name := "kk"
	nums := []int{1, 2, 3}

	func(name string, nums []int) {
		// 1.
		fmt.Println(name, nums) //
		name = "silence"
		nums = []int{1, 2}
		// 2.
		fmt.Println(name, nums) //
	}(name, nums)

	// 3.
	fmt.Println(name, nums) //
}
