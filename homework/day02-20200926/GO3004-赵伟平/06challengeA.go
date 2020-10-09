package main

import "fmt"

func insertionSort(old []int) {
	n := len(old)

	for i := 1; i < n; i++ {
		for j := i; j > 0 && old[j] < old[j-1]; j-- {
			old[j], old[j-1] = old[j-1], old[j]
		}
	}
}
func main() {
	nums := []int{108, 107, 105, 109, 103, 102}
	insertionSort(nums)
	fmt.Println(nums)
}
