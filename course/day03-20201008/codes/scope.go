package main

import "fmt"

func main() {
	name := "kk"
	nums := []int{}
	func() {
		//1.
		fmt.Println(name, nums) // kk, []
		name = "silence"
		nums = []int{1, 2, 3}
		//2.
		fmt.Println(name, nums) // silence, [1, 2, 3]
	}()
	//3.
	fmt.Println(name, nums) // silence, [1, 2, 3]

}
