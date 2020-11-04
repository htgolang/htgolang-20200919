package main

import "fmt"

var (
	slice = []int{108, 107, 105, 109, 103, 102}
)
//习题三，将最大的数移到最右边
func SortMax(this []int) {
	for i := 0; i < len(this) - 1; i++ {
		if this[i] > this[i + 1] {
			this[i + 1], this[i] = this[i], this[i + 1]
		}
	}
}

func main() {
	SortMax(slice)
	fmt.Println(slice)
}