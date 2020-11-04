package main

import "fmt"

var (
	slice = []int{108, 107, 105, 109, 103, 102}
)
//习题四，将最大的数移到最右边，并将第二大的数移到倒数第二位
func SortMax2(this []int) {
	//最大数
	for i := 0; i < len(this) - 1; i++ {
		if this[i] > this[i + 1] {
			this[i + 1], this[i] = this[i], this[i + 1]
		}
	}
	//第二大数
	for i := 0; i < len(this) - 2; i++ {
		if this[i] > this[i + 1] {
			this[i + 1], this[i] = this[i], this[i + 1]
		}
	}
}

func main() {
	SortMax2(slice)
	fmt.Println(slice)
}