package main

import "fmt"

func main() {
	name := "kk"
	nums := []int{1, 2, 3}

	func(pname string, pnums []int) {
		// 1.
		fmt.Println(pname, pnums) // kk, [1, 2, 3]
		pname = "silence"
		pnums[0] = 100
		// 2.
		fmt.Println(pname, pnums) // silence [100, 2, 3]
	}(name, nums)
	/*
		pname := name
		pname = "silence"

		pnums := nums
		pnums[0] = 100
	*/

	// 3.
	fmt.Println(name, nums) // kk, [100, 2, 3]
}
